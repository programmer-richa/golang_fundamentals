/*
Package main prints the reverse of array.
*/

package main

import "fmt"

// Entry point of the program
func main() {
	// Array to store 5 values
	var arr = [5]int{15, 10, 5, 20, 1}
	arr = reverse(arr)
	fmt.Printf("Reverse of array: %v", arr)
}

func reverse(arr [5]int) [5]int {
	size := len(arr)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
