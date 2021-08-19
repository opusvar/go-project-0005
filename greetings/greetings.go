package greetings

import (
	"errors"
	"fmt"
)

// return a greeting with an emeded name in in a message

func Hello(name string) (string, error) {
    if name == ""	 {
		return "", errors.New("empty name")
	}


	message := fmt.Sprintf("Hi %v. Welcome!", name)
	return message, nil
}