package main

import "fmt"

func main() {
	greeting := "Hello World!"
	// converting string to []byte
	byteSlice := []byte(greeting)
	fmt.Printf("Value and type of byteSlice is %v and %T", byteSlice, byteSlice)
}
