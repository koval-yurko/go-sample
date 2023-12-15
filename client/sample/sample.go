package sample

import (
	"fmt"

	"example.com/client/utils"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {

	utils.Utils_test()

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
