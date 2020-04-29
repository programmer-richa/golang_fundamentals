/*
Package main stores details of a student in a struct.
*/

package main

import "fmt"

// UnderGraduate stores detail of a student
type UnderGraduate struct {
	rollno int
	name   string
	marks  []float64
	total  float64
}

// SetRollNo assigns rollno
func (stu *UnderGraduate) SetRollNo(rollno int) {
	stu.rollno = rollno
}

// SetName assigns name
func (stu *UnderGraduate) SetName(name string) {
	stu.name = name
}

// SetMarks assigns marks
func (stu *UnderGraduate) SetMarks(marks []float64) {
	stu.marks = marks
	total := 0.0
	for _, v := range stu.marks {
		total += v
	}
	stu.total = total
}

// RollNo returns student rollno
func (stu UnderGraduate) RollNo() int {
	return stu.rollno
}

// Name returns student name
func (stu UnderGraduate) Name() string {
	return stu.name
}

// Marks returns student marks
func (stu UnderGraduate) Marks() []float64 {
	return stu.marks
}

// Total returns student total
func (stu *UnderGraduate) Total() float64 {
	return stu.total
}

// Input associates Input() method with UnderGraduate
func (stu *UnderGraduate) Input() {
	rollno, name, subjects := 0, "", 0
	fmt.Print("Enter student rollno :")
	fmt.Scan(&rollno)
	fmt.Print("Enter student name :")
	fmt.Scan(&name)
	fmt.Print("Enter number of subjects :")
	fmt.Scan(&subjects)
	marks := make([]float64, subjects)
	for i := range marks {
		fmt.Printf("Enter marks in subject %d: ", i+1)
		fmt.Scan(&marks[i])
	}
	stu.SetName(name)
	stu.SetRollNo(rollno)
	stu.SetMarks(marks)

}

// Output associates Input() method with UnderGraduate
func (stu UnderGraduate) Output() {
	fmt.Printf("Roll No: %d\n", stu.RollNo())
	fmt.Printf("Name: %s\n", stu.Name())
	fmt.Printf("Marks: %v\n", stu.Marks())
	fmt.Printf("Total: %f\n", stu.Total())

}

// Entry point of the program
func main() {
	// Stores student count
	number := 0
	fmt.Print("Enter  number of students: ")
	fmt.Scan(&number)
	// slice to store data of students
	students := make([]UnderGraduate, number)
	// loop over students to store student data
	for i := 0; i < number; i++ {
		fmt.Printf("Enter details of student %d\n", i+1)
		students[i] = UnderGraduate{}
		students[i].Input()
	}
	// loop over students to retrieve student data
	for i, s := range students {
		fmt.Printf("Details of student %d\n", i+1)
		s.Output()
	}

}
