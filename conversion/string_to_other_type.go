package main

import (
	"log"
	"strconv"
)

// StrToInt converts a string to int
func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panicf("Failed to convert string to int: %s", err)
	}
	return i
}

// StrToFloat64 converts a string to float64
func StrToFloat64(s string) float64 {
	bitSize := 64
	f, err := strconv.ParseFloat(s, bitSize)
	if err != nil {
		log.Panicf("Failed to convert string to float: %s", err)
	}
	return f
}

// StrToBool converts a string to bool
func StrToBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		log.Panicf("Failed to convert string to bool: %s", err)
	}
	return b
}
