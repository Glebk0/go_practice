package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main() {
	alex := person{firstName: "Alex",lastName: "Anderson"}
	fmt.Println(alex)
	var alex2 person
	alex2.firstName = "Alexander"
	fmt.Printf("%+v", alex2)
}