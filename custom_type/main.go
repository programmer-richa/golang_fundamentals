// Package main demonstrates working of custom type.
package main

import "fmt"

// Entry point of program.
func main() {
	deck := NewDeck()

	deck.Shuffle()
	deck.Print()
	hand, remainingCards := Deal(deck, 4)
	hand.Print()
	remainingCards.Print()

	// Saving a deck to a file
	writeErr := deck.SaveToFile("myCard")
	if writeErr == nil {

		fmt.Println("File written successfully.")
	} else {
		fmt.Println(writeErr.Error())
		fmt.Println("Error occured while writing file.")
	}
	newDeck := NewDeckFromFile("myCard")
	newDeck.Print()

}
