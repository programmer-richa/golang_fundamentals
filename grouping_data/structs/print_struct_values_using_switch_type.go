/*
Package main prints details of struct.
*/

package main

import "fmt"

type human interface {
	show()
}

type teacher struct {
	per person
	id  int
}

// attaching show() method with teacher
// this will also make teacher of type human too
func (tea teacher) show() {
	fmt.Printf("Employee ID : %d\n", tea.id)
	// calling show method of person
	tea.per.show()
}

type person struct {
	name, address string
	age           int
}

// attaching show() method with person
// this will also make person of type human too
func (per person) show() {
	fmt.Printf("Name : %s\n", per.name)
	fmt.Printf("Age : %d\n", per.age)
	fmt.Printf("Address : %s\n", per.address)
}

// function prints details of argument depending on its data type
func details(h human) {
	switch h.(type) {
	case person:
		fmt.Print("Person details:\n")
		h.show()
	case teacher:
		fmt.Print("Teacher details:\n")
		h.show()
	}
}

// Entry point of the program
func main() {
	// variable of type person
	per := person{
		name:    "Richa",
		address: "India",
		age:     26,
	}

	// variable of type teacher
	tec := teacher{
		id:  101,
		per: per,
	}

	fmt.Println("Person details using show method : ")
	per.show()

	fmt.Println("Teacher details using show method : ")
	tec.show()
	fmt.Println("\nDisplay details using type switch : ")
	details(per)
	details(tec)
}
