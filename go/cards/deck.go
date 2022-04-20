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
