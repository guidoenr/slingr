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
