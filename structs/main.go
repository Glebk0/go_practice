package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "alex@example.com",
			zipCode: 22230,
		},
	}
	fmt.Println(alex)
	var alex2 person
	alex2.firstName = "Alexander"
	fmt.Printf("%+v", alex2)
}
