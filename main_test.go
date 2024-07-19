package main

import (
	"os"
	"testing"
)

func TestWriteAndReadFile(t *testing.T) {
	path := "testfile.txt"
	content := "This is a test content."

	// Write to file
	if err := WriteToFile(path, content); err != nil {
		t.Fatalf("Failed to write to file: %v", err)
	}

	// Read from file
	readContent, err := ReadFromFile(path)
	if err != nil {
		t.Fatalf("Failed to read from file: %v", err)
	}

	// Verify the content
	if readContent != content {
		t.Errorf("Expected %q, got %q", content, readContent)
	}

	// Clean up
	if err := os.Remove(path); err != nil {
		t.Fatalf("Failed to remove test file: %v", err)
	}
}
