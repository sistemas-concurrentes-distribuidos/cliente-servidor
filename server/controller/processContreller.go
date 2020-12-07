package controller

import (
	"fmt"
	"net/http"
	"server/model"
	"server/services"
	s "server/services"
	util "server/utils"
	"strconv"

	"github.com/gorilla/mux"
)

//GetProcess handle dispatcher process endpoint
func GetProcess(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET: /process/dispatcher")
	process, value, err := s.GetActiveProcess()
	if err != nil {
		fmt.Println(err.Error())
		util.RespondWithError(w, http.StatusTeapot, err.Error())
		return
	}
	util.RespondWithJSON(w, http.StatusOK, model.AssingProcesResponse{IDProcess: process, Data: value})
}

//ReturnProcess handle collector process endpoint
func ReturnProcess(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET: /process/collector/{id}")
	vars := mux.Vars(r)
	processID, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if processID < 0 {
		util.RespondWithError(w, http.StatusBadRequest, "The ID cant be less than 0")
		return
	}
	services.ActiveProcess(processID)
	util.RespondWithJSON(w, http.StatusOK, map[string]string{"message:": "Task success"})
}
