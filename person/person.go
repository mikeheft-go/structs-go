package person

import "fmt"

type ContactInfo struct {
	Email   string
	ZipCode int32
}

type Person struct {
	FirstName string
	LastName  string
	ContactInfo
}

func (p Person) Print() {
	fmt.Printf("%+v", p)
}

func (pointer *Person) UpdateFirstName(newFirstName string) {
	(*pointer).FirstName = newFirstName
}
