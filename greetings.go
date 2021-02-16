package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// := is shorthand operator - declare + Initialize variable in one line
// alternative:
// 	var message string
// 	message = fmt.Sprintf("Hi %v. Welcome!", name)
