package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}

}

func main() {

	jobs := make(chan int, 5)
	result := make(chan int, 5)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, result)
	}

	for i := 0; i < 5; i++ {
		jobs <- i
	}

	// close(jobs)

	go printResult(result)

}

func printResult(result chan int) {
	for res := range result {
		fmt.Println("[results]", res)
	}
}
