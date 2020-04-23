/*
Package main calculates sum of digits of a number.
*/
package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	sum := sumDigits(number)
	fmt.Printf("Sum of digits of %d is %d", number, sum)
}

func sumDigits(number int) (sum int) {
	for number > 0 {
		remainder := number % 10
		number /= 10
		sum += remainder
	}
	return sum
}
