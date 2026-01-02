package main

import (
	"bufio"
	"fmt"
	"os"
)

type FileIO struct {
	FileName string
	Content  string
}

type FileOp interface {
	Check(name string) bool
	Create(name string, content string) error
	Read(name string) (string, error)
}

func main() {

}

func (f *FileIO) Check(name string) bool {

	_, err := os.Stat(name)

	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true

}

func (f *FileIO) Write(name string, content string) error {
	file, err := os.OpenFile(
		name,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0664,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(content + "\n")
	if err != nil {
		return err
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	return nil
}

func (f *FileIO) Read(name string) error {

	file, err := os.Open(name)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}
