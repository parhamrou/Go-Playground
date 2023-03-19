package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main() {
	parham := person{
		firstName: "Parham",
		lastName: "Roufarshbaf",
		contact: contactInfo{
			email: "p_roufarshbaf@yahoo.com",
			zipCode: 0000,
		},
	}
	fmt.Println(parham)
	parhamPointer := &parham
	parhamPointer.updateFirstName("PARHAM")
	fmt.Println(parham)

	// In go, if we defined a function with receiver of type A, we can use both pointer to A or the original A variable to use that function
	parham.updateFirstName("Parham")
	fmt.Println(parham)
}

func (personPointer *person) updateFirstName(newFirstName string) {
	(*personPointer).firstName = newFirstName
}