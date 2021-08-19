package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set the properties, flags, and ways to capture where the code isn't working
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Tiger")
	// catch an error, print it to the console, and exit the program
    if err != nil {
		log.Fatal(err)
	}

    // if no error is detected, run code as expected.
	fmt.Println(message)
}