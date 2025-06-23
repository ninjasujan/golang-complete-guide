package main

import "fmt"

func bootstrapServer(servers <-chan string) {

	// for server := range servers {
	// 	fmt.Println(server)
	// }

}

func main() {

	servers := make(chan string, 3)

	go bootstrapServer(servers)

	for i := 0; i < 3; i++ {
		servers <- fmt.Sprintf("server-%d", i)
	}

	for server := range servers {
		fmt.Println(server)
	}

	// Wait for the goroutine to complete
	close(servers)

}
