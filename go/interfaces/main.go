package main

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
