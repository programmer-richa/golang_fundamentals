// Package main demonstrates working of a function to interchange the values of 2 numbers.
package main

import "fmt"

// Entry point of a program
func main() {
	a, b := 10, 20
	fmt.Printf("Before swapping value of a = %d and b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swapping value of a = %d and b = %d", a, b)
}

// swap function accepts 2 arguments of *int type and interchange their values
func swap(x *int, y *int) {
	*x, *y = *y, *x
}
