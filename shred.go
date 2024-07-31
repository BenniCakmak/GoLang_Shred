package main

import (
	"crypto/rand"
	"os"
)

// Shred overwrites the given file three times with random data and then deletes the file
// path: String to the file
func Shred(path string) error {
	// Get the file handle
	file, err := os.OpenFile(path, os.O_WRONLY, 0)
	if err != nil {
		return err
	}

	// Get fileInfo to work with arguments like filesize
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// Get the actual file size
	fileSize := fileInfo.Size()
	// Allocate memory with the file size size
	randomData := make([]byte, fileSize)

	for i := 0; i < 3; i++ {
		// Fill randomData with random data
		_, err := rand.Read(randomData)
		if err != nil {
			return err
		}
		// Write randomData to file with no offset
		_, err = file.WriteAt(randomData, 0)
		if err != nil {
			return err
		}
		// Force flush of data into the physical data of the file to avoid caching-mechanisms
		err = file.Sync()
		if err != nil {
			return err
		}
	}

	// Close the file handle
	err = file.Close()
	if err != nil {
		return err
	}

	// Delete the file
	err = os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}
