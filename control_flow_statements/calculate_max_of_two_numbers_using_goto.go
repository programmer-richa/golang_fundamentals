/*
Package main prints max of 2 numbers entered by the user.
*/
package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	var first, second int
	fmt.Print("Enter first number: ")
	fmt.Scan(&first)
	fmt.Print("Enter second number: ")
	fmt.Scan(&second)
	if first > second {
		goto greater
	} else if second > first {
		goto smaller
	} else {
		goto equal
	}

greater:
	fmt.Printf("%d is greater than %d\n", first, second)
	goto end
smaller:
	fmt.Printf("%d is smaller than %d\n", first, second)
	goto end
equal:
	fmt.Printf("%d is equal to %d\n", first, second)
end:
	fmt.Print("Program Terminated.")

}
