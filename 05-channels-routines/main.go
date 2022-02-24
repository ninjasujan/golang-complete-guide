package main

import (
	"fmt"
	"net/http"
	"time"
)

func main () {
	links := [] string {
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.in",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
		time.Sleep(2)
	} 

	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink (link string, c chan string) {
	 _, err := http.Get(link)
	if err != nil {
		fmt.Println("Error in link", err)
		c <- "Might be down.!"
	}
	fmt.Println(link, "is working fine")
	c <- link
}