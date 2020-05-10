package main

import (
	"fmt"
	"strings"
)

func main() {
	colors := []string{"red", "blue", "green"}
	separator := ","
	// Joining slice
	color := strings.Join(colors, separator)
	fmt.Print(color)
}
