package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Starting program...")

	createFile("test.txt")

	fmt.Println("Program finished")

}

func createFile(filePath string) {

	defer recoverPanic()

	_, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	panic("force panic")

}

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}
