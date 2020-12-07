package services

import (
	"client/provider"
	"fmt"
	"time"
)

//Process logic to print data of the process
func Process(processID int, data int) {
	for {
		fmt.Printf("id %d: %d \n", processID, data)
		data++
		time.Sleep(time.Millisecond * 500)
	}
}

//StartProcess start a gorutine with the proces
func StartProcess(processID int, data int) {
	go Process(processID, data)
}

//StopProcess stop process an return the id
//process to the server
func StopProcess(processID int) {
	//TODO: implement stop process
	provider.ReturnProcess(processID)
}
