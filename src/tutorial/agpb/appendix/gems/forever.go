package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//for {}

	//select {}
	//defer func() { select {} }()

	//defer func() { <-make(chan bool) }()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel
	//time for cleanup before exit
	fmt.Println("Adios!")
}
