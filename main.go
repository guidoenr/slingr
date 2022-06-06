// DUMMY MAIN FOR TESTING
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	err := checkGoogleAccountCredentials()
	fmt.Println(err)
}

// maybe add it to the DOCKER RUN?
// export GOOGLE_APPLICATION_CREDENTIALS=/home/guido/service-account.json

func checkGoogleAccountCredentials() error {
	GoogleApplicationCredentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	fmt.Println(GoogleApplicationCredentials)
	if GoogleApplicationCredentials == "" {
		return errors.New("not found GOOGLE_APPLICATION_CREDENTIALS var, please set it")
	}
	return nil
}
