package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int32
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName:   "Jim",
		lastName:    "Party",
		contactInfo: contactInfo{email: "jim@email.com", zipCode: 80801},
	}
	// Need to update pointer to get original object to change
	jimPointer := &jim

	jimPointer.updateFirstName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointer *person) updateFirstName(newFirstName string) {
	(*pointer).firstName = newFirstName
}
