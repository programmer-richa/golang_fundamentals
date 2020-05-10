// Package main demonstrates working of a function expression.
package main

import (
	"fmt"
	"strings"
)

// Entry point of a program
func main() {
	a, b := 10, 20
	table := add(a, b)   // Calling the function and assigning function to a variable
	fmt.Print(table(20)) // Calling function using function expression
}

// add function accepts 2 arguments of int type and calculates their total.
// add returns a function that returns the string representing table of sum of numbers upto the multiplier passed as an argument.
func add(x int, y int) func(upto int) string {
	total := x + y

	return func(upto int) string {
		result := fmt.Sprintf(" Table of %d\n", total)
		result += strings.Repeat("-", 20) + "\n"
		for i := 1; i <= upto; i++ {
			result += fmt.Sprintf("%d * %d = %d\n", total, i, total*i)
		}
		return result
	}
}
