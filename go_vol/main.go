package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Path within the container corresponding to the mounted volume
	filePath := "/test/hello.txt"

	// Open the file in write mode, create it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write a line to the file
	line := "Hello from inside the container to the volume\n"
	_, err = file.WriteString(line)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Message written to %s\n", filePath)
}