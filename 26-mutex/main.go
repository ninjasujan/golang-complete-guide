package main

import (
	"fmt"
	"sync"
)

type Account struct {
	balance int64
	mutex   sync.Mutex
}

func (ac *Account) Withdraw(amount int64) {

	ac.mutex.Lock()
	defer ac.mutex.Unlock()

	if ac.balance >= amount {
		ac.balance -= amount
	}

}

func main() {

	var wg sync.WaitGroup

	account := &Account{balance: 1000}

	for i := 0; i < 10; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			// Simulate account withdrawal
			account.Withdraw(1)
		}()

	}

	wg.Wait()
	fmt.Println("Final account balance:", account.balance)

}
