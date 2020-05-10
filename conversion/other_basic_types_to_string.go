package main

import (
	"fmt"
	"strconv"
)

// IntToStr converts an int to string
func IntToStr(i int) string {
	s := strconv.Itoa(i)
	return s
}

// FloatToStr converts an float64 to string
func FloatToStr(f float64) string {
	s := fmt.Sprintf("%v", f)
	return s
}

// BoolToStr converts an bool to string
func BoolToStr(b bool) string {
	s := fmt.Sprintf("%v", b)
	return s
}

func main() {
	s := "40"
	i := StrToInt(s)
	ns := IntToStr(i)
	var x int64 = int64(i)
	bs := "true"
	b := StrToBool(bs)
	fmt.Printf("Type and value of s : %T and %v respectively\n", s, s)
	fmt.Printf("Type and value of i : %T and %v respectively\n", i, i)
	fmt.Printf("Type and value of ns : %T and %v respectively\n", ns, ns)
	fmt.Printf("Type and value of x : %T and %v respectively\n", x, x)
	fmt.Printf("Type and value of bs : %T and %v respectively\n", bs, bs)
	fmt.Printf("Type and value of b : %T and %v respectively\n", b, b)
}
