/*
Example of a game called Unique Six. In this game, the
player will roll a die to obtain a number from 1 to 6 and continue rolling until all six unique
numbers have been generated.
*/
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// declaring a slice to store unique numbers generated randomly
	generatedNumbers := make([]int, 6)
	// set unique numbers generated and user attempts to 0
	counter, attempts := 0, 0
	for {
		// variable to store generated number
		number := 0
		// generate number between 0 to 6
		for number == 0 {
			number = rand.Intn(7)
		}
		// print the generated number on the screen
		fmt.Println("Number rolled:", number)
		attempts++
		if !contains(generatedNumbers, number) {
			generatedNumbers[counter] = number
			counter++
		}
		fmt.Println("generatedNumbers:", generatedNumbers)

		if counter == 6 {
			fmt.Println("Congratulations, you have taken", attempts, "rolls to get one of each number between 1 and 6.")
			break
		} else {
			fmt.Println("So far, you have: ", sortSlice(generatedNumbers)[:counter+1])
		}
	}

}

// contains tests weather a value exists in the slice
func contains(slice []int, item int) bool {
	set := make(map[int]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func sortSlice(slice []int) []int {
	// sort slice in ascending order
	sort.Ints(slice)
	// reverse the slice variable
	slice = reverse(slice)
	return slice
}

// count returns number of unique choices made by user
func count(slice []int) (counter int) {
	slice = sortSlice(slice)
	// count elements greater than 0
	// looping over values in the slice
	for _, i := range slice {
		if i == 0 {
			break
		}
		counter++
	}
	return counter
}

// reverse reverses the order of elements in a slice
func reverse(slice []int) []int {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
	return slice
}
