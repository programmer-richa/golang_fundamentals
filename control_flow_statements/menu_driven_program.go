/*
Package main that works according to the user choice.
*/

package main

import (
	"fmt"
)

const (
	_ = iota
	leap
	palindrome
	digits
	table
	exit
	leapLbl       = "Test for leap year"
	palindromeLbl = "Test for Palindrom"
	digitsLbl     = "Calculate Sum of Digits"
	tableLbl      = "Print Table"
	exitLbl       = "Exit"
)

func options() {
	fmt.Printf("%d. %s\n", leap, leapLbl)
	fmt.Printf("%d. %s\n", palindrome, palindromeLbl)
	fmt.Printf("%d. %s\n", digits, digitsLbl)
	fmt.Printf("%d. %s\n", table, tableLbl)
	fmt.Printf("%d. %s\n", exit, exitLbl)
	fmt.Print("Enter your choice: ")
}

// Entry point of the program
func main() {
	var choice, number int
	for choice != 5 {
		fmt.Print("Enter a number: ")
		fmt.Scan(&number)
		options()
		fmt.Scan(&choice)
		switch choice {
		case leap:
			if testLeap(number) {
				fmt.Printf("%d is a leap year.\n", number)
			} else {
				fmt.Printf("%d is not a leap year.\n", number)
			}
		case palindrome:
			if testPalindrome(fmt.Sprint(number)) {
				fmt.Printf("%d is a palindrome number\n", number)
			} else {
				fmt.Printf("%d is not a palindrome number\n", number)
			}
		case digits:
			sum := sumD(number)
			fmt.Printf("Sum of digits of %d is %d", number, sum)
		case table:
			printTable(number)
		case exit:
			fmt.Println("Exiting...")
		default:
			fmt.Println("Invalid Choice. Try Again...")
		}
	}

}

func printTable(number int) {
	for i := 1; i < 11; i++ {
		fmt.Printf("%d * %d = %d\n", number, i, number*i)
	}
}

func testPalindrome(number string) (status bool) {
	reverse := reverseS(number)
	return reverse == number
}

func reverseS(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func sumD(number int) (sum int) {
	for number > 0 {
		remainder := number % 10
		number /= 10
		sum += remainder
	}
	return sum
}

func testLeap(year int) (status bool) {

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
