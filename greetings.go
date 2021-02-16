package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// if no name is given, returns an error
	if name == "" {
		return "", errors.New("empty name")
	}
	// if name was received, return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
	// nil = no error
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
			"Hi, %v. Welcome!",
			"Great to see you, %v!",
			"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
// := is shorthand operator - declare + Initialize variable in one line
// alternative:
// 	var message string
// 	message = fmt.Sprintf("Hi %v. Welcome!", name)

// go function can return multiple values