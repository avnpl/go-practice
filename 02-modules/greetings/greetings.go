package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error){
	// If no name, return an error with a message.
	if name == ""{
		return "", errors.New("Empty Name")
	}

    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Heil Hitler, %v!",
	}

	// Pick the current time as seed value for random number generation
	rand.Seed(time.Now().UnixNano())

	// Return a randomly selected message format by specifying
    // a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}