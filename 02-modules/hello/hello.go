package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

	names := []string{"John", "Paul", "George", "Ringo"}

    // Get a greeting message and print it.
    messages, err := greetings.Hellos(names)

	// If error, exit with error.
	if err != nil {
		log.Fatal(err)
	}

	// If no error, print message.
    fmt.Println(messages)
}
