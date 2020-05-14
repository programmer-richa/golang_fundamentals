package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Create a file and write string to it.
func writeTxtToFile() {
	fileName := "firstFile.txt"
	// Truncates if file already exists, be careful!
	file := CreateFile(fileName)
	// Make sure to close the file when file operation completes.
	defer CloseFile(file)
	defer ReadFile(fileName)
	WriteString(file, "This is my first file handling program in golang.")
}

// Returns a []Student
func getStudentSlice() (students []Student) {
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

	students = []Student{stu1, stu2}
	return students
}

// It converts []Student to JSON format and stores it to a variable. Then this variable reference is used to write in the file.
func jsonMarshakAndUnMarshal() {

	students := getStudentSlice()
	jsonFmt, err := json.Marshal(students)
	if err != nil {
		log.Fatalf("Failed converting student slice to json: %s", err)
	}
	fileName := "students.json"
	// Truncates if file already exists, be careful!
	file := CreateFile(fileName)
	// Make sure to close the file when file operation completes.
	defer CloseFile(file)
	defer ReadFile(fileName)
	WriteByte(fileName, jsonFmt)
}

// It converts []Student to JSON format and write in the file directly.
func encodeJSONToFile() {
	students := getStudentSlice()
	// Name of the file to work on.
	fileName := "student_data.json"
	// Create a new file if not exists and open in read write mode
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	// As file has Write method attached to it, it is of type Writer interface
	// Returning encoder pointer to file
	encoder := json.NewEncoder(file)
	//Ensure file is closed after encoded data is pushed to it.
	defer file.Close()
	//Encoding []Student and writing encoded data to the file.
	err = encoder.Encode(students)
	if err != nil {
		log.Fatal(err)
	}
}

// It reads JSON data in the file directly and converts stores into []student.
func decodeJSONToFile() {
	students := getStudentSlice()
	// Name of the file to work on.
	fileName := "student_data.json"
	// Create a new file if not exists and open in read write mode
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	// As file has Read method attached to it, it is of type Reader interface
	// Returning decoder pointer to file
	decoder := json.NewDecoder(file)
	//Ensure file is closed after decoded data is pushed to the slice.
	defer file.Close()
	//Decoding encoded data from the file to []Student.
	err = decoder.Decode(&students)
	if err != nil {
		log.Fatal(err)
	}
	for i, s := range students {
		fmt.Println("Student", i+1, ": ", s)
	}
}

func main() {
	writeTxtToFile()
	jsonMarshakAndUnMarshal()
	encodeJSONToFile()
	decodeJSONToFile()
}
