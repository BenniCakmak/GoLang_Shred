package main

import (
	"os"
	"testing"
)

// Testing if the file gets deleted at the end of Shred()
func TestShred(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up

	// Write some data to the temporary file
	if _, err := tmpfile.Write([]byte("This is a test")); err != nil {
		t.Fatal(err)
	}
	// Close the file handle
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Call the Shred function
	err = Shred(tmpfile.Name())
	if err != nil {
		t.Errorf("Shred() failed: %v", err)
	}

	// Check if the file has been deleted
	if _, err := os.Stat(tmpfile.Name()); !os.IsNotExist(err) {
		t.Errorf("File was not deleted")
	}
}

// Negative test - try to Shred a non existing file
func TestShredNonExistentFile(t *testing.T) {
	// Path to a non-existent file
	nonExistentFilePath := "non_existent_file.txt"

	// Ensure the file does not exist
	if _, err := os.Stat(nonExistentFilePath); !os.IsNotExist(err) {
		t.Fatalf("File %s unexpectedly exists", nonExistentFilePath)
	}

	// Call the Shred function
	err := Shred(nonExistentFilePath)
	if err == nil {
		t.Errorf("Expected error when shredding non-existent file, but got nil")
	}
}
