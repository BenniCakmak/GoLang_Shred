package main

import (
	"os"
	"fmt"
)

func main() {
	// Check for correct arguments
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <file_path>\n", os.Args[0])
		os.Exit(1)
	}
	
	path := os.Args[1]
	err  := Shred(path)
	if err != nil {
		panic(err)
	}
}
