// Package commands is where all bot commands go and being managed
package commands

import (
	"fmt"

	"github.com/diamondburned/arikawa/v2/bot"
)

type Bot struct {
	Ctx *bot.Context
}

// FailedCommand returns an error message if there was a problem with and process execution
func FailedCommand(command string, err error) (string, error) {
	// print error
	fmt.Println("command: "+command, err)

	return FailedMessage("There was a problem while trying to register, if the problem persists, please contact an admin and try again :slight_smile:.", err)
}

// FailedMessage is the message send on error or failed command
func FailedMessage(message string, err error) (string, error) {
	return "", fmt.Errorf(message)
}
