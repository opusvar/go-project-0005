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
    // returns a map that associates eaech of the named people with a specific greeting
	func Hellos(name []string) (map[string]string, error) {
		// A map to associate names with messages
		messages := make(map[string]string)
		// loop through the received slice of names, calling Hello to get a message
		// for each name
		for _, name := range names {
			message, err := Hello(name)
			if err != nil {
				return nil, err
			}
			// In map, associate the retrieved message with the name
			messages[name] = message
		}
		return messages, nil
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