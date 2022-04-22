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

	/* VALUE TYPES vs REFERENCE TYPES
	Value types: here, we got to use pointers to be able
	to change their value in a function
	[int, float, string, bool, structs]

	Reference Types: and here, we don't have to worry
	about pointers
	[slices ,maps, channels, pointers, functions ]

	*/

	//------------------ MAPS -------------------- //
	/*
	   map:
	   	key -> value
	   	key -> value

	   we could think that a Map in Go it's very similar to exactly
	   a Map in Python ;)
	*/

	colors := map[string]string{
		"red":  "ff0000",
		"blue": "0000ff",
	}
	// we can also define an empty map as:
	// colors := make(map[string]string)
	colors["white"] = "ffffff" //adding a key to the map
	fmt.Println(colors)

	// delete the element
	delete(colors, "white")

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
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
