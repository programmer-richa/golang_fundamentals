// Package main demonstrates working of a function to calculate sum of integers.
package main

import "fmt"

// Entry point of a program
func main() {
	a, b := 10, 20
	total := sum(a, b)
	fmt.Printf("Sum of %d and %d is %d", a, b, total)
}

// sum function accepts 2 arguments of int type and returns their total
func sum(x int, y int) (total int) {
	return x + y
}
