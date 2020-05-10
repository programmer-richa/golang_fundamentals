// Package main demonstrates working of json unmarshal to transform json to struct
package main

import (
	"encoding/json"
	"fmt"
)

//Entry point of a program
func main() {
	// declaring json variable of type byte slice
	jsonFmt := []byte(`[{"Rollno":101,"StudentName":"Richa","StudentMarks":{"English":90,"Maths":100,"Science":95}},{"Rollno":102,"StudentName":"Ritu","StudentMarks":{"English":70,"Maths":80,"Science":75}}]`)
	// creating slice of type Student
	students := []Student{}
	err := json.Unmarshal(jsonFmt, &students)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("JSON Data")
		fmt.Println(students)
		fmt.Println("Name of student at index 1", students[1].Name())
	}
}
