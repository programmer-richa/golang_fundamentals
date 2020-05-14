package main

// Student stores detail of a student
type Student struct {
	Roll     int
	StuName  string
	StuMarks []float64
	StuTotal float64
}

// SetRollNo assigns rollno
func (stu *Student) SetRollNo(rollno int) {
	stu.Roll = rollno
}

// SetName assigns name
func (stu *Student) SetName(name string) {
	stu.StuName = name
}

// SetMarks assigns marks
func (stu *Student) SetMarks(marks []float64) {
	stu.StuMarks = marks
}

// RollNo returns student rollno
func (stu Student) RollNo() int {
	return stu.Roll
}

// Name returns student name
func (stu Student) Name() string {
	return stu.StuName
}

// Marks returns student marks
func (stu Student) Marks() []float64 {
	return stu.StuMarks
}

// Total returns student total
func (stu *Student) Total() float64 {
	total := 0.0
	for _, v := range stu.StuMarks {
		total += v
	}
	stu.StuTotal = total
	return stu.StuTotal
}

// ByName defines methods listed in Interface interface in sort package.
// It enables to sort []Student by Name.
type ByName []Student

// Len is the number of alphabets in the name.
func (b ByName) Len() int {
	return len(b)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (b ByName) Less(i int, j int) bool {
	return b[i].Name() < b[j].Name()
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (b ByName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// ByTotal defines methods listed in Interface interface in sort package.
// It enables to sort []Student by Total Marks.
type ByTotal []Student

// Len is the number of alphabets in the name.
func (b ByTotal) Len() int {
	return len(b)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (b ByTotal) Less(i int, j int) bool {
	return b[i].Total() < b[j].Total()
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (b ByTotal) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
