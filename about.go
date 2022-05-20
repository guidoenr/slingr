// package = project = workspace
/*
a package can have a lot of files inside, like main.go, helper.go and support.go
but every file should has the 'package main' at the beginning,
why the 'main' name?
in go there are two types of packages

Executable: generates a file that we can run (like this one)
Reusable: code used as 'helpers' good place to put reusable logic, if we
are coding some kind of libraries or maybe a code to share with our friends
we gonna use the 'reusable' packages, not using the 'package main' at the top
something like 'package guido-tools' or whatelse

go build -> compiles a bunch of go source code files, and creates a executable 'main' file in the dir
and it can be runned with $ ./main
go run -> compiles and executes on or two fiels
go fmt -> fomrats all thhe code in each file in the current directory
go install -> compiles and installs a package [like maven]
go get -> downloads the raw source code [like maven]

*/
package main //so if the package is exectuable, there will be a 'package main' at the top of the file

import (
	"fmt" // a library to import debuggin and printing funcs
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hi there!")
}

/*
it's like a kind of method for the deck structure
this is a receiver, with the purpose of
" any variable of type 'deck' now gets access to the print method"
d: is a copy (always use d, or the first letter of the word)
deck: the reference to the type deck
*/
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// return a list of cards : deck
// we don't add any receiver in this function because
// we're not working with an existing deck, what we are doing right here
// its basically return a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Cloud"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// use '_' when we don't want to use the index variable
	// never use only a one variable
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// with (deck, deck) we are telling to go to return 2 type of values
// two decks
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/*
["red", "yellow", "blue"] -> "red, yellow, blue"
*/
func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

// note that the 'error' type its placed here because the
// writeFile function of GO also returns an error, this will be
// very usefull if we want to handle exceptions
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)

	if err != nil { // if err != nill means that there wasn't an error
		fmt.Println("Error:", err)
		os.Exit(1) // exit the program with -1 value
	}

	// because we want a string
	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}

func (d deck) shuffle() {

	// using a seed that changes by the current dateTime, so the seed
	// will be different everytime that we start our program
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source) // r = a Rand type

	for index := range d {
		// this will generate a random number between 0 and the deck's length
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index] // a simple double assignement for swap

	}
}


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

import "fmt"

// we are only creating this as an interface because we want to add
// some methods

type user struct {
	name string
}

// image this, as if we are creating an interface an adding methods
// notice that you could now create a "bot" variable, because this is an
// interface
type bot interface {
	getGreeting() string
	whoAmI() string
	respondToUser(u user) string
}

type englishBot struct {
}

type spanishBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	fmt.Println(eb.whoAmI())
	fmt.Println(eb.getGreeting())
	fmt.Println(eb.respondToUser(user{"guido"}))

	fmt.Println("-----------")

	fmt.Println(sb.whoAmI())
	fmt.Println(sb.getGreeting())
	fmt.Println(sb.respondToUser(user{"guido"}))

}

// note: WE DON'T need to put 'eb englishBot' in this case,
// because we aren't using the param, so, this looks really
// like a method
func printGreeting(b bot) {
	fmt.Println(b.getGreeting()) // POLYMORFISM
}

// --------------- SPANISH BOT ---------------

// don't be scared, you dont need to link the interface
// and the bots,
func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (spanishBot) whoAmI() string {
	return "Soy el bot español"
}

func (spanishBot) respondToUser(u user) string {
	return "Hola " + u.name
}

// --------------- ENGLISH BOT ---------------

func (englishBot) getGreeting() string {
	return "Hi There!"
}

func (englishBot) whoAmI() string {
	return "I am the english bot"
}

func (englishBot) respondToUser(u user) string {
	return "Hello " + u.name
}

// -READER INTERFACE:
// the reader interface can receive several type of data
// and then create an output date of type []byte that anyone
// can work with it

func main() {
	//  var card string = "Ace of Spades"  this is STATIC, this will never contain a number
	/*
		this line is the same that the one above, i'ts also static, but with the difference
		that GO automaticlly infers the variable type (string)
		and this line, also works for initializing and assigning a new value to a variable
	*/
	//card := "Ace of spades" // MOST USED
	// the ':=' its only used when we are declaring and defining a new variable in one line
	// if we use only the '=' we are making an assigment to an existing variable
	//card = "Five of Diamonds"
	//new_card := newCard()

	//fmt.Println(new_card)
	//fmt.Println(card)

	// ----- ARRAYS - SLICE ------- //
	// ARRAY: fixed length list of things
	// SLICE: an array that can grow or shrink (every element must be of the same type)

	// --------------------------------- using the deck.go

	//cards[0:2].print() // cards[startIndexIncluding : upToNotIncluding] , this also works like cards[:2]
	// and the reverse will be [2:] this returns everything in the right of 2

	cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	//hand.print()
	// fmt.Println("---------------remaining deck: ")
	// remainingDeck.print()

	// ------ BYTE SLICE: a string represented in a computer [72 105 32 116 104 101 114 101 33]
	//byteSlice := []byte("Hi There") // this is the way to cast a string to a byte Slice
	//fmt.Println(byteSlice)

	string_cards := cards.toString()
	fmt.Println(string_cards)

	cards.saveToFile("cards.dat")

	// if we put an invalid filename, that will throws the error with Exit(1)
	deckFromFile := newDeckFromFile("cards.dat")
	fmt.Println(deckFromFile)

	// shuffle - random
	kards2 := newDeck()
	kards2.shuffle()
	kards2.print()
}

// always put the data type in return-type func
func newCard() string {
	return "Five of diamonds"
}

/*
GO is NOT and oriented object language
- its static
*/