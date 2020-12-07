package router

import (
	"net/http"
	c "server/controller"

	"github.com/gorilla/mux"
)

//InitRouter config the router and add all the endpoint
func InitRouter() *mux.Router {
	//Create the router
	r := mux.NewRouter()

	//Add endpoint
	r.HandleFunc("/process/collector/{id}", c.ReturnProcess).Methods(http.MethodGet)
	r.HandleFunc("/process/dispatcher", c.GetProcess).Methods(http.MethodGet)

	return r
}
