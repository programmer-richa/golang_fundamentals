// Package main demonstrates working of json marshal to transform struct to json
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Entry point of a program
func main() {
	// declaring first variable of type Student
	s1 := Student{}
	mark1 := map[string]float64{"Maths": 100, "Science": 95, "English": 90}
	s1.SetRollNo(101)
	s1.SetName("Richa")
	s1.SetMarks(mark1)

	// declaring second variable of type Student
	s2 := Student{}
	mark2 := map[string]float64{"Maths": 80, "Science": 75, "English": 70}
	s2.SetRollNo(102)
	s2.SetName("Ritu")
	s2.SetMarks(mark2)

	// creating slice of Student
	students := []Student{s1, s2}
	jsonFmt, err := json.Marshal(students)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		//os.Stdout.Write(b)
		fmt.Printf("Type of jsonFmt : %T\n", jsonFmt)
		os.Stdout.Write(jsonFmt)
	}
}
