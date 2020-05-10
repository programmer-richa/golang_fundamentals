package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ReadFile reads the data of the file and prints it.
func ReadFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("Failed reading data from file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)
	fmt.Printf("\nError: %v", err)
}

// CreateFile creates the file with given name and returns file reference.
func CreateFile(fileName string) *os.File {
	// Truncates if file already exists, be careful!
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}
	return file
}

// WriteFile writes data to the file.
func WriteFile(file *os.File, data string) {
	// Truncates if file already exists, be careful!
	len, err := file.WriteString(data)

	if err != nil {
		log.Fatalf("Failed writing to file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len)
	fmt.Printf("\nFile Name: %s", file.Name())
}

//CloseFile closes the file
func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatalf("Failed closing file: %s", err)
	}
}
