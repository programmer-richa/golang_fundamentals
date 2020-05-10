package main

import (
	"fmt"
	"strings"
)

func main() {
	byteCards := []byte{65, 99, 101, 32, 111, 102, 32, 67, 108, 117, 98, 115, 44, 84, 119, 111, 32, 111, 102, 32, 67, 108, 117, 98, 115, 44, 84, 104, 114, 101, 101, 32, 111, 102, 32, 67, 108, 117, 98, 115, 44, 70, 111, 117, 114, 32, 111, 102, 32, 67, 108, 117, 98, 115, 44, 70, 105, 118, 101, 32, 111, 102, 32, 67, 108, 117, 98, 115, 44, 83, 105, 120, 32, 111, 102, 32, 67, 108, 117, 98, 115, 44, 83, 101, 118, 101, 110, 32, 111, 102, 32, 67, 108, 117, 98, 115, 44, 69, 105, 103, 104, 116, 32, 111, 102, 32, 67, 108, 117, 98, 115, 44, 78, 105, 110, 101, 32, 111, 102, 32, 67, 108, 117, 98, 115, 44,
		84, 101, 110, 32, 111, 102, 32, 67, 108, 117, 98, 115, 44, 74, 97, 99, 107, 32, 111, 102, 32, 67, 108, 117, 98, 115, 44, 81, 117, 101, 101, 110, 32, 111, 102, 32, 67, 108, 117, 98, 115, 44, 75, 105, 110, 103, 32, 111, 102, 32, 67, 108, 117, 98, 115, 44, 65, 99, 101, 32, 111, 102, 32, 83, 112, 97, 100, 101, 115, 44, 84, 119, 111, 32, 111, 102, 32, 83, 112, 97, 100, 101, 115, 44, 84,
		104, 114, 101, 101, 32, 111, 102, 32, 83, 112, 97, 100, 101, 115, 44, 70, 111, 117, 114, 32, 111, 102, 32, 83, 112, 97, 100, 101, 115, 44, 70, 105, 118, 101, 32, 111, 102, 32, 83, 112, 97, 100, 101, 115, 44, 83, 105, 120, 32, 111, 102, 32, 83, 112, 97, 100, 101, 115, 44, 83, 101, 118, 101, 110, 32, 111, 102, 32, 83, 112, 97, 100, 101, 115, 44, 69, 105, 103, 104, 116, 32, 111, 102, 32, 83, 112, 97, 100, 101, 115, 44, 78, 105, 110, 101, 32, 111, 102, 32, 83, 112, 97, 100, 101, 115, 44, 84, 101, 110, 32, 111, 102, 32, 83, 112, 97, 100, 101, 115, 44, 74, 97, 99, 107, 32, 111, 102, 32, 83, 112, 97, 100, 101, 115, 44, 81, 117, 101, 101, 110, 32, 111, 102, 32, 83, 112, 97, 100, 101, 115, 44, 75, 105, 110, 103, 32, 111, 102, 32, 83, 112, 97, 100, 101, 115, 44, 65,
		99, 101, 32, 111, 102, 32, 68, 105, 97, 109, 111, 110, 100, 115, 44, 84, 119, 111, 32, 111, 102, 32, 68, 105, 97, 109, 111, 110, 100, 115, 44, 84, 104, 114, 101, 101, 32, 111, 102, 32, 68, 105, 97, 109, 111, 110, 100, 115, 44, 70, 111, 117, 114, 32, 111, 102, 32, 68, 105, 97, 109, 111, 110, 100, 115, 44, 70, 105, 118, 101, 32, 111, 102, 32, 68, 105, 97, 109, 111, 110, 100, 115, 44, 83, 105, 120, 32, 111, 102, 32, 68, 105, 97, 109, 111, 110, 100, 115, 44, 83, 101, 118, 101, 110, 32, 111, 102, 32, 68, 105, 97, 109, 111, 110, 100, 115, 44, 69, 105, 103, 104, 116, 32, 111, 102, 32, 68, 105, 97, 109, 111, 110, 100, 115, 44, 78, 105, 110, 101, 32, 111, 102, 32, 68, 105, 97, 109, 111, 110, 100, 115, 44, 84, 101, 110, 32, 111, 102, 32, 68, 105, 97, 109, 111, 110, 100, 115, 44, 74, 97, 99, 107, 32, 111, 102, 32, 68, 105, 97, 109, 111, 110, 100, 115, 44, 81, 117, 101, 101, 110, 32, 111, 102, 32, 68, 105, 97, 109, 111, 110, 100, 115, 44, 75, 105, 110, 103, 32, 111, 102, 32, 68, 105, 97, 109, 111, 110, 100, 115, 44, 65, 99, 101, 32, 111, 102, 32, 72, 101, 97, 114, 116, 115, 44, 84, 119, 111, 32, 111, 102, 32, 72, 101, 97, 114, 116, 115, 44, 84, 104, 114, 101, 101, 32, 111, 102, 32, 72, 101, 97, 114, 116, 115, 44, 70, 111, 117, 114, 32, 111, 102, 32, 72, 101, 97, 114, 116, 115, 44, 70, 105, 118, 101, 32, 111, 102, 32, 72, 101, 97,
		114, 116, 115, 44, 83, 105, 120, 32, 111, 102, 32, 72, 101, 97, 114, 116, 115, 44, 83, 101, 118, 101, 110, 32, 111, 102, 32, 72, 101, 97, 114, 116, 115, 44, 69, 105, 103, 104, 116, 32, 111, 102, 32, 72, 101, 97, 114, 116, 115, 44, 78, 105, 110, 101, 32, 111, 102, 32, 72, 101, 97, 114, 116, 115, 44, 84, 101, 110, 32, 111, 102, 32, 72, 101, 97, 114, 116, 115, 44, 74, 97, 99, 107, 32, 111, 102, 32, 72, 101, 97, 114, 116, 115, 44, 81, 117, 101, 101, 110, 32, 111, 102, 32, 72, 101, 97, 114, 116, 115, 44, 75, 105, 110, 103, 32, 111, 102, 32, 72, 101, 97, 114, 116, 115}
	card := string(byteCards)
	cards := strings.Split(card, ",")
	fmt.Print(cards)
}