package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, res chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		res <- j
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {
	numJobs := 5
	numWorkers := 3
	var wg sync.WaitGroup

	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, result, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		fmt.Println("Result - ", <-result)
	}

	wg.Wait()
	close(result)

}
