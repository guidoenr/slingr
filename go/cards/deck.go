package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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
