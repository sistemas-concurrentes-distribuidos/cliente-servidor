package main

import (
	"fmt"
	"net/http"
	"server/router"
	s "server/services"
)

func main() {

	//Config the router
	r := router.InitRouter()

	//Active all process
	s.StartAllProcess()

	//Start server
	if err := http.ListenAndServe(":8000", r); err != nil {
		fmt.Printf("Error al inciar el servidor. %+v", err)
		return
	}
}
