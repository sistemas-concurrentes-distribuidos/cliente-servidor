package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//callback function to be executed at exit intercept
type callback func()

//InterceptExit intercept exit and execute some collaback
func InterceptExit(callback callback) {
	//catch contrl+c to return process to the sever
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		callback()
		fmt.Printf("\nProceso entregado\n")
		os.Exit(1)
	}()
}
