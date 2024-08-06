package main

import "structs-go/person"

func main() {
	jim := person.Person{
		FirstName:   "Jim",
		LastName:    "Party",
		ContactInfo: person.ContactInfo{Email: "jim@email.com", ZipCode: 80801},
	}
	// Need to update pointer to get original object to change
	jim.UpdateFirstName("John")
	jim.Print()
}
