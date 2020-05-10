package main

/*
Student type holds values regarding roll no, name and marks of a student
Here all the attributes are defined using initial capital letters will be accessible outside the package
The attribute to be exported with JSON marshal should be an exported member.
*/
type Student struct {
	// RollNo stores student rollno
	Rollno int
	// Name stores student name
	StudentName string
	// Marks stores student marks in different subjects
	StudentMarks map[string]float64
}

// SetRollNo assigns rollno value to the student
func (s *Student) SetRollNo(rollno int) {
	s.Rollno = rollno
}

// SetName assigns name value to the student
func (s *Student) SetName(name string) {
	s.StudentName = name
}

// SetMarks assigns marks value to the student
func (s *Student) SetMarks(marks map[string]float64) {
	s.StudentMarks = marks
}

// RollNo returns rollno value to the student
func (s Student) RollNo() (rollno int) {
	return s.Rollno
}

// Name returns name value to the student
func (s Student) Name() (name string) {
	return s.StudentName
}

// Marks returns marks value to the student
func (s Student) Marks() (marks map[string]float64) {
	return s.StudentMarks
}
