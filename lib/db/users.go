package db

import (
	"github.com/World-of-Cryptopups/cordy/lib"
	"github.com/deta/deta-go/service/base"
)

// GetAllUsers attempts to get all users from db.
func (dc *DetaClient) GetAllUsers() ([]*lib.User, error) {
	var results []*lib.User
	var page []*lib.User

	// fetch input
	i := &base.FetchInput{
		Q:    base.Query{},
		Dest: &page,
	}

	// fetch items
	lKey, err := dc.DB.Fetch(i)
	if err != nil {
		return nil, err
	}

	results = append(results, page...)

	for lKey != "" {
		i.LastKey = lKey

		// fetch items
		lKey, err = dc.DB.Fetch(i)
		if err != nil {
			return nil, err
		}

		results = append(results, page...)
	}

	return results, nil
}
