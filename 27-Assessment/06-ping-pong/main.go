package main

import (
	"fmt"
	"time"
)

func PingGo(ping, pong chan struct{}) {

	for {
		<-ping
		fmt.Println("[PING]")
		time.Sleep(time.Second * 2)
		pong <- struct{}{}
	}

}

func PongGo(ping, pong chan struct{}) {

	for {
		<-pong
		fmt.Println("[PONG]")
		time.Sleep(time.Second * 2)
		ping <- struct{}{}
	}
}

func main() {

	pingCh := make(chan struct{})
	pongCh := make(chan struct{})

	go PingGo(pingCh, pongCh)
	go PongGo(pingCh, pongCh)

	pingCh <- struct{}{}

	// select {}

}
