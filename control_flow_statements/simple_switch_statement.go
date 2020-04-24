/*
Package main that works according to the choice.
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

	switch choice := 4; choice {
	case 1:
		if testLeap1(number) {
			fmt.Printf("%d is a leap year.\n", number)
		} else {
			fmt.Printf("%d is not a leap year.\n", number)
		}
	case 2:
		if testPalindrome1(fmt.Sprint(number)) {
			fmt.Printf("%d is a palindrome number\n", number)
		} else {
			fmt.Printf("%d is not a palindrome number\n", number)
		}
	case 3:
		sum := sumD1(number)
		fmt.Printf("Sum of digits of %d is %d", number, sum)
	case 4:
		printTable1(number)
	case 5:
		fmt.Println("Exiting...")
	default:
		fmt.Println("Invalid Choice. Try Again...")
	}

}

func printTable1(number int) {
	for i := 1; i < 11; i++ {
		fmt.Printf("%d * %d = %d\n", number, i, number*i)
	}
}

func testPalindrome1(number string) (status bool) {
	reverse := reverseS1(number)
	return reverse == number
}

func reverseS1(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func sumD1(number int) (sum int) {
	for number > 0 {
		remainder := number % 10
		number /= 10
		sum += remainder
	}
	return sum
}

func testLeap1(year int) (status bool) {

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
