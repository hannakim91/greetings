package greetings

import (
	"errors"
	"fmt"
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
}

// := is shorthand operator - declare + Initialize variable in one line
// alternative:
// 	var message string
// 	message = fmt.Sprintf("Hi %v. Welcome!", name)
