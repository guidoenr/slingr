package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	validString()
}

func validString() {
	bad_chars := []string{"â", "€", "™"}
	for i := range bad_chars {
		fmt.Printf("valid string: %v -> %v \n", bad_chars[i], utf8.ValidString(bad_chars[i]))
	}

}
