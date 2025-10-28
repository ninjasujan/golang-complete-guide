package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	workerCount = 3
	jobCount    = 5
)

func Worker(id int, jobs <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		result <- j * 2
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {

	jobs := make(chan int, jobCount)
	result := make(chan int, jobCount)

	var wg sync.WaitGroup

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go Worker(i, jobs, result, &wg)
	}

	for j := 1; j <= jobCount; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= jobCount; i++ {
		fmt.Println("Result - ", <-result)
	}

	wg.Wait()
	close(result)

}
