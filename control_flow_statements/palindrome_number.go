/*
Package main tests whether the number entered by the user is palindrome.
*/
package main

import (
	"fmt"
)

// Entry point of the program
func main() {
	var number string
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	if isPalindrome(number) {
		fmt.Printf("%s is a palindrome number", number)
	} else {
		fmt.Printf("%s is not a palindrome number", number)
	}
}

func isPalindrome(number string) (status bool) {
	reverse := reverseString(number)
	fmt.Println("Reverse : ", reverse)
	return reverse == number
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
