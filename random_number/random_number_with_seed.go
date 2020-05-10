package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNumber := 6
	// Fetch time now
	timeNow := time.Now().UnixNano()
	//create seed value of type Source in rand package
	source := rand.NewSource(timeNow)
	// Get Rand struct variable reference
	r := rand.New(source)
	for i := 0; i < 15; i++ {

		//Generates random number in different sequence every time the program is executed
		num := r.Intn(maxNumber)
		fmt.Println("Random number generated: ", num)
	}

}
