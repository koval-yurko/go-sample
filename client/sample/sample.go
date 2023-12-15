package sample

import (
	"fmt"

	"github.com/koval-yurko/go-sample/client/utils"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {

	utils.Utils_test()

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome v0.0.7!", name)
	return message
}
