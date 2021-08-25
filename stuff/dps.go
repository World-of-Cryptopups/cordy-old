package stuff

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/World-of-Cryptopups/cordy/lib"
	"github.com/World-of-Cryptopups/cordy/utils"
)

// FetchDPS is a fetcher to call the endpoint and save to db.
func FetchDPS(user lib.UserDPSUser, wallet string) (lib.UserDPSInfo, error) {
	r, err := utils.PostFetcher(user, os.Getenv("DPS_FETCH")+wallet)
	if err != nil {
		fmt.Printf("\n error in posting data in FetchDPS, error: %v", err)
		return lib.UserDPSInfo{}, err
	}

	var data lib.UserDPSInfo
	if err := json.Unmarshal(r, &data); err != nil {
		fmt.Printf("\n error in unmarshalling data in FetchDPS, error: %v", err)
		return lib.UserDPSInfo{}, err
	}

	return data, nil
}

// Get the DPS of a certain discordId user.
func GetDPS(id string) (lib.UserDPSInfo, error) {
	r, err := utils.Fetcher(os.Getenv("DPS_GET") + id)
	if err != nil {
		fmt.Printf("\n error in posting data in GetDPS, error: %v", err)
		return lib.UserDPSInfo{}, err
	}

	var data lib.UserDPSInfo
	if err := json.Unmarshal(r, &data); err != nil {
		fmt.Printf("\n error in unmarshalling data in GetDPS, error: %v", err)
		return lib.UserDPSInfo{}, err
	}

	return data, nil
}
