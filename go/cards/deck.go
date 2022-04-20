package main

import "fmt"

// create a new type of 'deck'
// which is a slice of strings
type deck []string // it's like defining an object, a deck is a slice of strings now
type entero int    // this is another example

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
func newDeck() {
	
}
