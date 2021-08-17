package fauna

import (
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
