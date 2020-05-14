// Package main demonstrates working of custom type.
package main

import (
	"fmt"
	"sort"
)

// Entry point of program.
func main() {
	students := []Student{}
	students = append(students, Student{Roll: 103, StuName: "Richa", StuMarks: []float64{90, 88, 79}})
	students = append(students, Student{Roll: 101, StuName: "Sahil", StuMarks: []float64{90, 85, 70}})
	students = append(students, Student{Roll: 102, StuName: "Aman", StuMarks: []float64{40, 58, 69}})
	fmt.Println("Students Before Sorting:", students)
	sort.Sort(ByName(students))
	fmt.Println("Students After Sorting By Name :", students)
	sort.Sort(ByTotal(students))
	fmt.Println("Students After Sorting By Total :", students)

}
