package main

import (
	"client/provider"
	s "client/services"
	"client/utils"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Client...")

	response, err := provider.GetProcess()
	if err != nil {
		log.Printf("error: %v \n", err.Error())
	}

	if response.StatusCode == 418 {
		fmt.Printf("\nNo hay procesos disponibes\n")
		return
	}

	s.StartProcess(response.IDProcess, response.Data)

	//catch control+c to return proces to the server
	utils.InterceptExit(func() { s.StopProcess(response.IDProcess) })

	//PAUSE
	fmt.Scanln()

	//return the process to the server
	s.StopProcess(response.IDProcess)
	fmt.Printf("\nProceso entregado\n")
}
