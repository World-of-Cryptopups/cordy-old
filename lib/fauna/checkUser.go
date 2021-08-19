package fauna

import (
	e "github.com/World-of-Cryptopups/roleroll-new/lib/errors"
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

// CheckUser checks if the user exists.
func CheckUser(discordId string) (bool, error) {
	client := Client()

	_registered, err := client.Query(f.Exists(f.MatchTerm(f.Index("userByDiscordId"), discordId)))
	if err != nil {
		return false, err
	}
	var registered bool
	_registered.Get(&registered)

	return registered, nil
}

// IsUserRegistered checks if user is registered. It wraps calculations.
func IsUserRegistered(discordID string) (interface{}, error) {
	_registered, err := CheckUser(discordID)
	if err != nil {
		return e.FailedCommand("check if user is registered", err)
	}
	if !_registered {
		return e.FailedMessage("You are not registered! You can register by sending `>register {your-token}`.", err)
	}

	return "", nil
}
