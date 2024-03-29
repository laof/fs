package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/laof/fs/utils"

	"github.com/julienschmidt/httprouter"
)

type Sizer interface {
	Size() int64
}

type Data struct {
	Success bool  `json:"success"`
	Size    int64 `json:"size"`
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func smartName(dirPath string, filename string, i int) string {

	path := filepath.Base(filename)
	ext := filepath.Ext(path)
	pre := path[0 : len(path)-len(ext)]

	if i == 0 {
		path = pre + ext
	} else {
		path = pre + strconv.Itoa(i) + ext
	}

	if fileExists(filepath.Join(dirPath, path)) {
		return smartName(dirPath, filename, i+1)
	} else {
		return path
	}
}

func UploadFiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// dir
	file, handler, err := r.FormFile("upload")
	dir := r.FormValue("dir")
	if err != nil {
		http.Error(w, err.Error(), 500)
		fail := Data{false, 0}
		f := utils.JsonData(fail)
		w.Write(f)
		return
	}

	defer file.Close()
	ff := smartName(dir, handler.Filename, 0)
	savePath := filepath.Join(dir, ff)
	fmt.Println("Upload => ", savePath)
	f, _ := os.Create(savePath)

	defer f.Close()
	io.Copy(f, file)

	data := Data{true, handler.Size}
	s := utils.JsonData(data)
	w.Write(s)
}
