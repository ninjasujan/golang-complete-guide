package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(ctx context.Context, id int, wg *sync.WaitGroup) {

	defer wg.Done()

	select {
	case <-time.After(time.Second * 10):
		fmt.Println("[worker executed]", id)

	case <-ctx.Done():
		fmt.Println("[worker cancelled due to timeout]")
	}

}

func main() {

	worker := 3
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go Worker(ctx, i, &wg)

	}

	wg.Wait()

	fmt.Println("[Execution completed]")

}
