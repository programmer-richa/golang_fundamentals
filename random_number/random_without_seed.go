package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxNumber := 6
	for i := 0; i < 15; i++ {
		//Generates random number in same sequence every time the program is executed
		num := rand.Intn(maxNumber + 1)
		fmt.Println("Random number generated: ", num)
	}

}
