package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	fmt.Println(quote.Go())

	names := []string{"Ray", "Amy", "Rachel"}
	messages, err := greetings.Hellos(names)
	// If error returned; print to console and exit
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
