package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// return a greeting with an emeded name in in a message

func Hello(name string) (string, error) {
    if name == ""	 {
		return "", errors.New("empty name")
	}

	// create a message using a random format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}
	// set the inital values for variables usedin the function
	func init() {
		rand.Seed(time.Now().UnixNano())
	}
    
	// randomFormat returns one sete of greetings, selected at random

	func randomFormat() string {
		// slice of message formats
		formats := []string {
			"Hi, %v.  Welcome!",
			"Great to see you, %v!",
			"Hail, %v!  Well met!",
		}
		// returns a random message format by specficying a random index for the slice of formats.
		return formats[rand.Intn(len(formats))]
	}