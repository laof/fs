package controllers

import (
	"files-server/models"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func IndexFiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// http.ServeFile(w, r, filepath.Join("./", path))
	http.ServeFile(w, r, filepath.Join("./", "index.html"))

}

var smap map[string]string

func lazyload() {
	if len(smap) == 0 {
		smap = models.GetStaticMapping()
	}
}

type NotFoundHttpServe struct {
}

func (h *NotFoundHttpServe) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	lazyload()

	path := r.URL.Path

	if smap[path] == "" {

		http.ServeFile(w, r, filepath.Join("./", path))

	} else {

		if strings.Contains(path, ".css") {
			w.Header().Add("content-type", "text/css; charset=utf-8")
		} else {
			w.Header().Add("content-type", "application/javascript; charset=utf-8")
		}

		w.Write([]byte(smap[path]))

	}
}

func Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// _, err := r.Cookie("fs")
	// if err != nil {
	// 	ck := &http.Cookie{
	// 		Name:     "fs",
	// 		Value:    strconv.FormatInt(time.Now().UnixNano(), 10),
	// 		Path:     "/",
	// 		Secure:   true,
	// 		HttpOnly: true,
	// 		MaxAge:   0}
	// 	http.SetCookie(w, ck)
	// }
	lazyload()
	w.Write([]byte(smap["/main.html"]))
}
