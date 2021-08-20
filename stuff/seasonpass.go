package stuff

import (
	"encoding/json"
	"os"

	"github.com/World-of-Cryptopups/cordy/utils"
)

// GetSeasonOnePass is a getter for getting the Season One Pass DPS.
func GetSeasonOnePass(wallet string) (SeasonPass, error) {
	r, err := utils.Fetcher(os.Getenv("SEASONPASS_ONE_GET") + wallet)
	if err != nil {
		return SeasonPass{}, err
	}

	var data SeasonPass
	if err := json.Unmarshal(r, &data); err != nil {
		return SeasonPass{}, err
	}

	return data, nil
}

// ConfirmSeasonOnePass is a easonOne Pass Verifier
func ConfirmSeasonOnePass(wallet string) (SeasonPassVerify, error) {
	r, err := utils.Fetcher(os.Getenv("SEASONPASS_ONE_VERIFY") + wallet)
	if err != nil {
		return SeasonPassVerify{}, err
	}

	var data SeasonPassVerify
	if err := json.Unmarshal(r, &data); err != nil {
		return SeasonPassVerify{}, err
	}

	return data, nil
}
