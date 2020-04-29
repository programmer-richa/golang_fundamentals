/*
Package main stores details of a student in a struct.
*/

package main

import "fmt"

// Student stores detail of a student
type Student struct {
	rollno int
	name   string
	marks  []float64
	total  float64
}

// SetRollNo assigns rollno
func (stu *Student) SetRollNo(rollno int) {
	stu.rollno = rollno
}

// SetName assigns name
func (stu *Student) SetName(name string) {
	stu.name = name
}

// SetMarks assigns marks
func (stu *Student) SetMarks(marks []float64) {
	stu.marks = marks
}

// RollNo returns student rollno
func (stu Student) RollNo() int {
	return stu.rollno
}

// Name returns student name
func (stu Student) Name() string {
	return stu.name
}

// Marks returns student marks
func (stu Student) Marks() []float64 {
	return stu.marks
}

// Total returns student total
func (stu *Student) Total() float64 {
	total := 0.0
	for _, v := range stu.marks {
		total += v
	}
	stu.total = total
	return stu.total
}

// Entry point of the program
func main() {
	// Creating and initializing values to student 1
	stu1 := Student{}
	stu1.SetRollNo(101)
	stu1.SetName("Richa")
	stu1.SetMarks([]float64{90, 85, 100})

	// Creating and initializing values to student 2
	stu2 := Student{}
	stu2.SetRollNo(102)
	stu2.SetName("Ritu")
	stu2.SetMarks([]float64{80, 95, 70})

	if stu2.Total() > stu1.Total() {
		fmt.Printf("%s scored more marks than %s\n", stu2.Name(), stu1.Name())

	} else {
		fmt.Printf("%s scored more marks than %s\n", stu1.Name(), stu2.Name())
	}

	fmt.Printf("Record of student 1: %v\n", stu1)
	fmt.Printf("Record of student 2: %v\n", stu2)

}
