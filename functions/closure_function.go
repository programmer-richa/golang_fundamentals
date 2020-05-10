// Package main demonstrates working of a closure function to generate a sequence.
package main

import "fmt"

// Entry point of a program
func main() {
	firstSequence := sequence()
	fmt.Printf("Number Generated %d\n", firstSequence())
	fmt.Printf("Number Generated %d\n", firstSequence())

	secondSequence := sequence()
	fmt.Printf("Number Generated %d\n", secondSequence())
	fmt.Printf("Number Generated %d\n", secondSequence())
}

// sum function accepts 2 arguments of int type and returns their total
func sequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
