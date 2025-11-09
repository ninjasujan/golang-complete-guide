package main

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context) {

	fmt.Println("[worker started]")

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Processing complete")
		// Simulate long processing
	case <-ctx.Done():
		// Handle cancellation
		fmt.Println("Processing cancelled:", ctx.Err())
		return
	}
	fmt.Println("[worker ended]")

}

func main() {

	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	go func() {
		process(ctx)
	}()

	<-ctx.Done()

	select {
	case <-ctx.Done():
		fmt.Println("Main context done:", ctx.Err())
	default:
		fmt.Println("Main context still active")
	}

}
