package main

import (
	"fmt"
	"os"
	"path/filepath"
)

const dirPath = "./notes/"

func main() {
	fmt.Println("[File Operation In Go...]")

	// Check and create directory if it doesn't exist
	if err := checkDirectory(dirPath); err != nil {
		fmt.Printf("Error checking directory: %v\n", err)

	}

	// Create a test file
	err := createFile(dirPath, "todo-3.txt", "ToDo List 1...")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)

	}

	todoList, err := ListToDo(dirPath)
	if err != nil {
		fmt.Printf("Error listing todo list: %v\n", err)

	}

	fmt.Println(todoList)

}

// checkDirectory checks if the directory exists and creates it if it doesn't
func checkDirectory(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
		}
		fmt.Printf("Directory created: %s\n", dirPath)
	} else if err != nil {
		return fmt.Errorf("error checking directory %s: %w", dirPath, err)
	}
	return nil
}

// createFile creates a new file in the specified directory
func createFile(dirPath string, fileName string, note string) error {
	filePath := filepath.Join(dirPath, fileName)

	// Check if file already exists
	if _, err := os.Stat(filePath); err == nil {
		return fmt.Errorf("file already exists: %s", filePath)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("error checking file %s: %w", filePath, err)
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	if _, err := file.Stat(); err != nil {
		return fmt.Errorf("failed to get file info: %w", err)
	}

	if _, err := file.WriteString(note); err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filePath, err)
	}

	return nil

}

func ListToDo(dirPath string) ([]string, error) {

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory %s: %w", dirPath, err)
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("no files found in directory %s", dirPath)
	}

	todoList := make([]string, len(files))

	for index, file := range files {
		if !file.IsDir() {
			fileData, err := os.ReadFile(filepath.Join(dirPath, file.Name()))
			if err != nil {
				return nil, fmt.Errorf("failed to read file %s: %w", file.Name(), err)
			}
			todoList[index] = string(fileData)
		}
	}

	return todoList, nil

}
