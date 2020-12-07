package main

import (
	c "cliente-servidor/controller"
	s "cliente-servidor/services"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/process/{id}", c.ReturnProcess).Methods(http.MethodPost)
	r.HandleFunc("/process", c.GetProcess).Methods(http.MethodGet)

	//Active all process
	s.StartAllProcess()

	if err := http.ListenAndServe(":8000", r); err != nil {
		fmt.Printf("Error al inciar el servidor. %+v", err)
		return
	}
}
