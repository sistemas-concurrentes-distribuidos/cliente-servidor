package services

import (
	"errors"
	"fmt"
	"time"
)

var active bool = false
var contador uint64 = 0

var statusProcess = []bool{true, true, true, true, true}

func Proceso() {
	for {
		for index, active := range statusProcess {
			if active {
				fmt.Printf("id %d: %d \n", index, contador)
			}
		}
		fmt.Println("---------------------------")
		contador += 1
		time.Sleep(time.Millisecond * 500)
	}
}

func StartAllProcess() {
	if !active {
		active = true
		go Proceso()
	}
}

func StopProcess(processId uint64) {
	statusProcess[processId] = false
}

func ActiveProcess(processId uint64) {
	statusProcess[processId] = true
}

//return proccesId, actual value of process and error
func GetActiveProcess() (uint64, uint64, error) {
	for i, v := range statusProcess {
		if v {
			StopProcess(uint64(i))
			return uint64(i), contador, nil
		}
	}
	return 0, 0, errors.New("no processes availables")
}
