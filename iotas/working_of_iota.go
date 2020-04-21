// Package main specifies that the variables, constants and functions
// declared in this folder belongs to the main namespace
package main

import (
	"fmt"
	"strings"
)

// Declaring constants using a const block that automatically assigns value of constant declared above to the uninitialized constant in the block.
// A constant block groups the constants.
const (
	// It is mandatory to assign a value to the first declared constant in a block.
	// The value in iota is incremented by 1 for each constant declaration.
	// The value of iota is reset to zero for each constant block.
	zeroSet1  = iota // iota = 0
	oneSet1          // iota = 1
	twoSet1          // iota = 2
	threeSet1        // iota = 3
	fourSet1         // iota = 4
)

// Declaring constants using a const block that automatically assigns value of constant declared above to the uninitialized constant in the block.
// A constant block groups the constants.
const (
	// It is mandatory to assign a value to the first declared constant in a block.
	// The value in iota is incremented by 1 for each constant declaration.
	// The value of iota is reset to zero for each constant block.
	// Using _ operator to skip the value
	_         = iota // iota = 0
	oneSet2          // iota = 1
	twoSet2          // iota = 2
	threeSet2        // iota = 3
	fourSet2         // iota = 4
)

// Declaring constants using a const block that automatically assigns value of constant declared above to the uninitialized constant in the block.
// A constant block groups the constants.
const (
	// It is mandatory to assign a value to the first declared constant in a block.
	// The value in iota is incremented by 1 for each constant declaration, even if some other value is explicitly assigned to a constant.
	// The value of iota is reset to zero for each constant block.
	// Using _ operator to skip the value
	_        = iota // iota = 0
	oneSet3         // iota = 1
	twoSet3         // iota = 2
	other    = 50   // iota = 3
	fourSet3 = iota // iota = 4
)

// Entry point of a program
func main() {
	fmt.Println("Types and values stored in constant block Set 1")
	// The below code uses the Repeat function of strings package to repeat '-' for 100 times.
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("Type and value of package level untyped constant 'zeroSet1' is '%T' and '%v' respectively.\n", zeroSet1, zeroSet1)
	fmt.Printf("Type and value of package level untyped constant 'oneSet1' is '%T' and '%v' respectively.\n", oneSet1, oneSet1)
	fmt.Printf("Type and value of package level untyped constant 'twoSet1' is '%T' and '%v' respectively.\n", twoSet1, twoSet1)
	fmt.Printf("Type and value of package level untyped constant 'threeSet1' is '%T' and '%v' respectively.\n", threeSet1, threeSet1)
	fmt.Printf("Type and value of package level untyped constant 'fourSet1' is '%T' and '%v' respectively.\n", fourSet1, fourSet1)

	fmt.Println(strings.Repeat("-", 100))
	fmt.Println(strings.Repeat("-", 100))

	fmt.Println("Types and values stored in constant block Set 2")
	// The below code uses the Repeat function of strings package to repeat '-' for 100 times.
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("Type and value of package level untyped constant 'oneSet2' is '%T' and '%v' respectively.\n", oneSet2, oneSet2)
	fmt.Printf("Type and value of package level untyped constant 'twoSet2' is '%T' and '%v' respectively.\n", twoSet2, twoSet2)
	fmt.Printf("Type and value of package level untyped constant 'threeSet2' is '%T' and '%v' respectively.\n", threeSet2, threeSet2)
	fmt.Printf("Type and value of package level untyped constant 'fourSet2' is '%T' and '%v' respectively.\n", fourSet2, fourSet2)

	fmt.Println(strings.Repeat("-", 100))
	fmt.Println(strings.Repeat("-", 100))

	fmt.Println("Types and values stored in constant block Set 3")
	// The below code uses the Repeat function of strings package to repeat '-' for 100 times.
	fmt.Println(strings.Repeat("-", 100))
	fmt.Printf("Type and value of package level untyped constant 'oneSet3' is '%T' and '%v' respectively.\n", oneSet3, oneSet3)
	fmt.Printf("Type and value of package level untyped constant 'twoSet3' is '%T' and '%v' respectively.\n", twoSet3, twoSet3)
	fmt.Printf("Type and value of package level untyped constant 'other' is '%T' and '%v' respectively.\n", other, other)
	fmt.Printf("Type and value of package level untyped constant 'fourSet3' is '%T' and '%v' respectively.\n", fourSet3, fourSet3)

}
