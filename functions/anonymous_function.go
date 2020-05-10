// Package main demonstrates working of an anonymous function.
package main

import "fmt"

// Entry point of a program
func main() {
	a, b := 10, 20
	// anonymous function accepts 2 arguments of int type and returns their total
	total := func (x int, y int) (total int) {
		return x + y
	} (a,b) // here (a,b) is calling the anonymous function by passing a and b as actual arguments
	fmt.Printf("Sum of %d and %d is %d", a, b, total)
}

