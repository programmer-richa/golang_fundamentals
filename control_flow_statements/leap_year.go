/*
Package main tests whether the year entered by the user is a leap year.
*/
package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	var year int
	fmt.Print("Enter a year: ")
	fmt.Scan(&year)
	if isLeap(year) {
		fmt.Printf("%d is a leap year.", year)
	} else {
		fmt.Printf("%d is not a leap year.", year)
	}
}
func isLeap(year int) (status bool) {

	// If a year is multiple of 400,
	// then it is a leap year
	if year%400 == 0 {
		return true
	}
	// Else If a year is multiple of 100,
	// then it is not a leap year
	if year%100 == 0 {
		return false
	}

	// Else If a year is multiple of 4,
	// then it is a leap year
	if year%4 == 0 {
		return true
	}
	return false
}
