package main

import "fmt"

type person struct {
	firstName string
	lastname  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	//axel := person{"Axel", "Rose"}
	//fmt.Println(axel.firstName, axel.firstName, axel.lastname)
	//fmt.Printf("%+v", axel) //this prints the format %+v
	//// updating the struct value
	//axel.lastname = "updatedLastName"
	//axel.firstName = "updatedFirstName"
	//fmt.Printf("%+v", axel)

	john := person{
		firstName: "John",
		lastname:  "Doe",
		contact: contactInfo{
			email:   "johndoe@gmail.com",
			zipCode: 1337,
		},
	}
	/* POINTERS:
	&variable = memory address of the value this variable it's pointing at
	*pointer = the value this memory address it's pointing at
	*/

	// this is the pointer to john
	johnPointer := &john
	johnPointer.updateName("newName")

	john.print()
}

// a receiver with a struct
func (p person) print() {
	fmt.Printf("%+v", p)
}

// a receiver to update the name. but with a pointer
// to the person
func (p *person) updateName(newName string) {
	(*p).firstName = newName
}
