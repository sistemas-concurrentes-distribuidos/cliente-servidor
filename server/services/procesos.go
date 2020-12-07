package services

import (
	"errors"
	"fmt"
	"time"
)

var active bool = false
var contador int = 0

var statusProcess = []bool{true, true, true, true, true}

func proceso() {
	for {
		for index, active := range statusProcess {
			if active {
				fmt.Printf("id %d: %d \n", index, contador)
			}
		}
		fmt.Println("---------------------------")
		contador++
		time.Sleep(time.Millisecond * 500)
	}
}

//StartAllProcess active the log of the process
func StartAllProcess() {
	if !active {
		active = true
		go proceso()
	}
}

//StopProcess stop the actual process
func StopProcess(processID int) {
	statusProcess[processID] = false
}

//ActiveProcess active the log of the process
//with the ID given
func ActiveProcess(processID int) {
	statusProcess[processID] = true
}

//GetActiveProcess get the firs active procces of the
//statusProcess slide
//return proccesId, actual value of process and error
func GetActiveProcess() (int, int, error) {
	for i, v := range statusProcess {
		if v {
			StopProcess(i)
			return i, contador, nil
		}
	}
	return 0, 0, errors.New("no processes availables")
}
