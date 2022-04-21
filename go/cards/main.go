package main

import "fmt"

// deckSize := 20
// variables can be initialized outside of a function, but cannot be assigned a variable
// the correct way will be deckSize
var deckSize int

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
	deckFromFile := newDeckFromFile("cardss.dat")
	fmt.Println(deckFromFile)
}

// always put the data type in return-type func
func newCard() string {
	return "Five of diamonds"
}

/*
GO is NOT and oriented object language
- its static
*/
