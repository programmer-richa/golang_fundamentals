// Package main specifies that the variables, constants and functions
// declared in this folder belongs to the main namespace
package main

// Importing required modules
import (
	"fmt"
)

// Entry point of a program
func main() {
	// Declaring variable with data type
	// Since there is no concept of unintialized variables in go, all the variables are assigned zero values.
	var byte1 byte // preferred type in file handling or working with streams

	var int1 int // preferred type for integer values
	var int2 int8
	var int3 int16
	var int4 int32
	var int5 int64

	var uint1 uint
	var uint2 uint8
	var uint3 uint16
	var uint4 uint32
	var uint5 uint64

	var float1 float32
	var float2 float64 // preferred data type for storing floating point

	var complex1 complex64
	var complex2 complex128

	var string1 string

	var bool1 bool
	// Calling Printf function available in fmt package for printing the output on console
	fmt.Printf("Value stored in variable int : %d\n ", int1)
	fmt.Printf("Value stored in variable int8 : %d\n ", int2)
	fmt.Printf("Value stored in variable int16 : %d\n ", int3)
	fmt.Printf("Value stored in variable int32 : %d\n ", int4)
	fmt.Printf("Value stored in variable int64 : %d\n ", int5)

	fmt.Printf("Value stored in variable uint : %d\n ", uint1)
	fmt.Printf("Value stored in variable uint8 : %d\n ", uint2)
	fmt.Printf("Value stored in variable uint16 : %d\n ", uint3)
	fmt.Printf("Value stored in variable uint32 : %d\n ", uint4)
	fmt.Printf("Value stored in variable uint64 : %d\n ", uint5)

	fmt.Printf("Value stored in variable float32 : %f\n ", float1)
	fmt.Printf("Value stored in variable float64 : %f\n ", float2)

	fmt.Printf("Value stored in variable complex64 : %g\n ", complex1)
	fmt.Printf("Value stored in variable complex128 : %g\n ", complex2)

	fmt.Printf("Value stored in variable string : %s\n ", string1)

	fmt.Printf("Value stored in variable bool : %t\n ", bool1)

	fmt.Printf("Value stored in variable bool : %b\n ", byte1)
}
