package errors

import (
	"errors"
	"fmt"
)

// FailedCommand returns an error message if there was a problem with and process execution
func FailedCommand(command string, err error) (string, error) {
	// print error
	fmt.Println("command: "+command, err)

	return FailedMessage("There was a problem while trying to execute the command, if the problem persists, please contact an admin and try again :slight_smile:.", err)
}

// FailedMessage is the message send on error or failed command
func FailedMessage(message string, err error) (string, error) {
	fmt.Println(err)

	return "", errors.New(message)
}

// error if not registered
func RegisterErr() (string, error) {
	return FailedMessage("You are not registered! You can register by sending `>register {your-token}`. Please get your token by signing in to https://www.worldofcryptopups.cf/ 😉", nil)
}
