package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("Worker received stop signal")
			return
		default:
			fmt.Println("Worker running...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)

	stop := make(chan struct{})

	go worker(stop)

	<-signChan // Block until signal received
	fmt.Println("[OS Exit]")
	close(stop)
	time.Sleep(500 * time.Millisecond) // Give worker time to print exit message
}
