package controller

import (
	"cliente-servidor/model"
	"cliente-servidor/services"
	s "cliente-servidor/services"
	util "cliente-servidor/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetProcess(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET: /process")
	process, value, err := s.GetActiveProcess()
	if err != nil {
		fmt.Println(err.Error())
		util.RespondWithError(w, http.StatusTeapot, err.Error())
		return
	}
	util.RespondWithJson(w, http.StatusOK, model.AssingProcesResponse{IdProcess: process, Data: value})
}

func ReturnProcess(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST: /process/{id}")
	vars := mux.Vars(r)
	processId, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if processId < 0 {
		util.RespondWithError(w, http.StatusBadRequest, "The ID cant be less than 0")
		return
	}
	services.ActiveProcess(uint64(processId))
	util.RespondWithJson(w, http.StatusOK, map[string]string{"message:": "Task success"})
}
