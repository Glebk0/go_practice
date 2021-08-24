package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@example.com",
			zipCode: 22230,
		},
	}
	fmt.Println(alex)
	var alex2 person
	alex2.firstName = "Alexander"
	// &variable - give the memory address of the value this variable is pointing at
	alex2Pointer := &alex2 // tunring value in address
	alex2Pointer.updateName("alexxx")
	alex2.printinfo()

	//short way of changing structs. Method receiver should still have pointer for that to work
	alex.updateName("newalex")
	alex.printinfo()
}

// *pointer - give the value this memory address is pointing at
// in this case *person in receiver is TYPE DESCRIPTION - it means function works with a pointer to a person
// in function body *pointerToPerson is an OPERATOR - it means we want to manipulate the value the pointer is referencing
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) printinfo() {
	fmt.Printf("%+v\n", p)
}
