package main

import (
	"fmt"
	"io"
	"os"
)

// WriteToFile writes the given content to a file at the specified path.
func WriteToFile(path string, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// ReadFromFile reads the content from a file at the specified path.
func ReadFromFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func main() {
	path := "example.txt"
	content := "Hello, Go file handling!"

	// Write content to file
	if err := WriteToFile(path, content); err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
	fmt.Println("Content written to file successfully.")

	// Read content from file
	readContent, err := ReadFromFile(path)
	if err != nil {
		fmt.Printf("Error reading from file: %v\n", err)
		return
	}
	fmt.Printf("Content read from file: %s\n", readContent)
}
