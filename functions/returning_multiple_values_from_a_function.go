// Package main demonstrates working of a function expression.
package main

import (
	"fmt"
	"strings"
)

// Entry point of a program
func main() {
	a, b := 10, 20
	total, table := plus(a, b) // Calling the function and assigning total and function to a variable
	fmt.Printf("Sum of %d and %d is %d\n", a, b, total)
	fmt.Print(table(20)) // Calling function using function expression
}

// plus function accepts 2 arguments of int type and calculates their total.
// plus returns total calculated and a function that returns the string representing table of sum of numbers upto the multiplier passed as an argument.
func plus(x int, y int) (int, func(upto int) string) {
	total := x + y

	return total, func(upto int) string {
		result := fmt.Sprintf(" Table of %d\n", total)
		result += strings.Repeat("-", 20) + "\n"
		for i := 1; i <= upto; i++ {
			result += fmt.Sprintf("%d * %d = %d\n", total, i, total*i)
		}
		return result
	}
}
