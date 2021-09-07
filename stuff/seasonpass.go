package stuff

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/World-of-Cryptopups/cordy/lib"
	"github.com/World-of-Cryptopups/cordy/utils"
)

// GetCurrentPass gets the current user's season pass.
func GetCurrentPass(wallet string) (lib.SeasonPassVerify, error) {
	r, err := utils.Fetcher(fmt.Sprintf("%s/seasonpass/one/%s", os.Getenv("CORDY_API"), wallet))
	if err != nil {
		return lib.SeasonPassVerify{}, err
	}

	var data lib.SeasonPassVerify
	if err := json.Unmarshal(r, &data); err != nil {
		return lib.SeasonPassVerify{}, err
	}

	return data, nil
}

// GetSeasonOnePass is a getter for getting the Season One Pass DPS.
func GetSeasonOnePass(wallet string) (lib.SeasonPass, error) {
	r, err := utils.Fetcher(os.Getenv("SEASONPASS_ONE_GET") + wallet)
	if err != nil {
		return lib.SeasonPass{}, err
	}

	var data lib.SeasonPass
	if err := json.Unmarshal(r, &data); err != nil {
		return lib.SeasonPass{}, err
	}

	return data, nil
}

// ConfirmSeasonOnePass is a easonOne Pass Verifier
func ConfirmSeasonOnePass(wallet string) (lib.SeasonPassVerify, error) {
	r, err := utils.Fetcher(os.Getenv("SEASONPASS_ONE_VERIFY") + wallet)
	if err != nil {
		return lib.SeasonPassVerify{}, err
	}

	var data lib.SeasonPassVerify
	if err := json.Unmarshal(r, &data); err != nil {
		return lib.SeasonPassVerify{}, err
	}

	return data, nil
}
