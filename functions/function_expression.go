// Package main demonstrates working of a function expression.
package main

import "fmt"

// Entry point of a program
func main() {
	a, b := 10, 20
	// anonymous function accepts 2 arguments of int type and returns their total
	sum := func(x int, y int) (total int) {
		return x + y
	}
	total := sum(a, b) // Calling the function by its expression name
	fmt.Printf("Sum of %d and %d is %d", a, b, total)
}
