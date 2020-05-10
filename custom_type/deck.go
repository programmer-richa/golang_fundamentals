package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Deck is a custom type which is created on the slice of string.
type Deck []string

// NewDeck returns a new deck of cards.
func NewDeck() (d Deck) {
	suits := []string{"Clubs", "Spades", "Diamonds", "Hearts"}
	numbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	// d = make([]string, 52)
	for _, suit := range suits {
		for _, number := range numbers {
			d = append(d, number+" of "+suit)
		}
	}
	return d
}

//Print prints all the values in the deck.
func (d Deck) Print() {
	for num, card := range d {
		fmt.Println(num+1, ". ", card)
	}
}

// Deal divides deck into two parts.
func Deal(d Deck, handSize int) (hand Deck, remainingCards Deck) {
	return d[:handSize], d[handSize:]
}

// ToString returns string representation of cards in deck separated by ','.
func (d Deck) ToString() string {
	// conveting deck to []string, and then joining the values in the slice to form a string.
	return strings.Join([]string(d), ",")
}

// ToByteSlice returns []byte representation of the deck.
func (d Deck) ToByteSlice() []byte {
	return []byte(d.ToString())
}

// SaveToFile saves the information in the deck to a file.
func (d Deck) SaveToFile(filename string) error {
	err := ioutil.WriteFile(filename, d.ToByteSlice(), 0755)
	return err
}

// DeckFromByteSlice creates a deck from a []byte.
func DeckFromByteSlice(byteCards []byte) (d Deck) {
	card := string(byteCards)
	cards := strings.Split(card, ",")
	return Deck(cards)
}

// NewDeckFromFile reads data from the file and creates a byte
func NewDeckFromFile(filename string) Deck {
	byteCards, err := ioutil.ReadFile(filename)
	if err == nil {
		return DeckFromByteSlice(byteCards)
	} else {
		// Error occured while reading file
		// Create new deck and return it.
		// or call os.Exit(1) to forcefully exit the program.
		return NewDeck()
	}
}

// Shuffle suffles the deck
func (d Deck) Shuffle() {
	maxNumber := len(d)
	// Fetch time now
	timeNow := time.Now().UnixNano()
	//create seed value of type Source in rand package
	source := rand.NewSource(timeNow)
	// Get Rand struct variable reference
	r := rand.New(source)
	for i := 0; i < 15; i++ {

		//Generates random number in different sequence every time the program is executed
		num := r.Intn(maxNumber)
		d[i], d[num] = d[num], d[i]
	}

}
