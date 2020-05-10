// Package main demonstrates working of a call back function expression.
package main

import (
	"fmt"
	"strings"
)

// Entry point of a program
func main() {
	a, b := 10, 20
	str := addition(table, a, b) // Passing table as a call back function
	fmt.Print(str + "\n")
	str = addition(print, a, b) // Passing print as a call back function
	fmt.Printf("Sum of %d and %d is %s\n", a, b, str)

}

// addition function accepts a function as the first argument that returns a string value.
// a series of int values are passed as arguments to the function
// addition calculates sum of int arguments and pass the total calculated to the function argument, and returns the result produced by that function.
func addition(f func(number int) string, x ...int) (result string) {
	total := 0
	for _, v := range x {
		total += v
	}
	return f(total)
}

func table(number int) string {
	result := fmt.Sprintf(" Table of %d\n", number)
	result += strings.Repeat("-", 20) + "\n"
	for i := 1; i <= 10; i++ {
		result += fmt.Sprintf("%d * %d = %d\n", number, i, number*i)
	}
	return result
}

func print(number int) string {
	result := fmt.Sprintf("%d\n", number)
	return result
}
