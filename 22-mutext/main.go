package main

import (
	"fmt"
	"sync"
)

var counter = 0

var mutex = sync.Mutex{}

func increment(wg *sync.WaitGroup) {
	mutex.Lock()
	counter++
	mutex.Unlock()
	wg.Done()
}

func main() {

	expCounter := 2000
	wg := sync.WaitGroup{}

	for i := 0; i < expCounter; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()

	fmt.Println(counter)

}
