package controllers

import (
	"net/http"

	"github.com/laof/fs/models"
	"github.com/laof/fs/utils"

	"github.com/julienschmidt/httprouter"
)

type HostData struct {
	Host     string `json:"host"`
	Shutdown bool   `json:"shutdown"`
}

func GetHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	data := HostData{models.GetHostAddress(), false}
	w.Write(utils.JsonData(data))
}
