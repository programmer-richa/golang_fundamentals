/*
Package main generates a table of a number.
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
	fmt.Printf("Table of %d\n", number)
	tableNumber(number)
}

func tableNumber(number int) {
	for i := 1; i < 11; i++ {
		fmt.Printf("%d * %d = %d\n", number, i, number*i)
	}
}
