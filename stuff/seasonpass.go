package stuff

import (
	"encoding/json"
	"os"

	"github.com/World-of-Cryptopups/cordy/lib"
	"github.com/World-of-Cryptopups/cordy/utils"
)

// GetCurrentPass gets the current user's season pass.
func GetCurrentPass(wallet string) (lib.SeasonPassVerify, error) {
	r, err := utils.Fetcher(os.Getenv("SEASONPASS_CURRENT_GET") + wallet)
	if err != nil {
		return lib.SeasonPassVerify{}, err
	}

	var data lib.SeasonPassVerify
	if err := json.Unmarshal(r, &data); err != nil {
		return lib.SeasonPassVerify{}, err
	}

	return data, nil
}
