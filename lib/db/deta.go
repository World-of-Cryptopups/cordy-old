package db

import (
	"os"

	"github.com/World-of-Cryptopups/cordy/lib"
	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)

type DetaClient struct {
	DB    *base.Base
	DPSDB *base.Base
	D     *deta.Deta
}

// Client creates / returns a new Deta base client.
func Client() (*DetaClient, error) {
	d, err := deta.New(deta.WithProjectKey(os.Getenv("DETA_PROJECT_KEY")))
	if err != nil {
		return nil, err
	}

	db, err := base.New(d, "WoC_DATA")
	if err != nil {
		return nil, err
	}

	dpsDB, err := base.New(d, "DPSInfos")
	if err != nil {
		return nil, err
	}

	return &DetaClient{
		D:     d,
		DB:    db,
		DPSDB: dpsDB,
	}, nil
}

// GetUser gets the user by its key. Key is the discordID
func (dc *DetaClient) GetUser(discordID string) (*lib.User, error) {
	var r *lib.User

	if err := dc.DB.Get(discordID, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// queryUser used for querying that returns a user
func (dc *DetaClient) queryUser(query base.Query) ([]*lib.User, error) {
	var r []*lib.User

	_, err := dc.DB.Fetch(&base.FetchInput{
		Q:     query,
		Dest:  &r,
		Limit: 1,
	})
	if err != nil {
		return nil, err
	}

	return r, nil
}

// UserExists checks if the user already has been registered or not.
func (dc *DetaClient) UserExists(discordID string) (bool, error) {
	users, err := dc.queryUser(base.Query{
		{"user.id": discordID},
	})

	// there was a problem, false but with error
	if err != nil {
		return false, err
	}

	if len(users) > 0 {
		return true, nil
	}

	return false, nil
}

// TokenExists checks if the token already exists in the db.
func (dc *DetaClient) TokenExists(token string) (bool, error) {
	users, err := dc.queryUser(base.Query{
		{"token": token},
	})

	// there was a problem, false but with error
	if err != nil {
		return false, err
	}

	if len(users) > 0 {
		return true, nil
	}

	return false, nil
}
