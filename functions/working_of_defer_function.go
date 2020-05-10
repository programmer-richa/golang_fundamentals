// Package main demonstrates working of a function to interchange the values of 2 numbers.
package main

import "fmt"

// Entry point of a program
func main() {
	defer message("Closed the file")
	defer message("Closing the file")
	defer message("Reading the file")
	message("Opening the file")
}

// message accepts an argument of string type and prints that string on the screen
func message(msg string) {
	fmt.Println(msg)
}
