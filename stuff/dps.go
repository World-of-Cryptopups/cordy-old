package stuff

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/World-of-Cryptopups/cordy/lib"
	"github.com/World-of-Cryptopups/cordy/utils"
)

type FetchDPSResponse struct {
	Success bool            `json:"success"`
	Data    lib.UserDPSInfo `json:"data"`
	Message string          `json:"message"`
}

// FetchDPS is a fetcher to call the endpoint and save to db.
func FetchDPS(user lib.UserDPSUser, wallet string) (lib.UserDPSInfo, error) {
	r, err := utils.PostFetcher(user, fmt.Sprintf("%s/dps/calculator/%s", os.Getenv("CORDY_API"), wallet))
	if err != nil {
		fmt.Printf("\n error in posting data in FetchDPS, error: %v", err)
		return lib.UserDPSInfo{}, err
	}

	var data FetchDPSResponse
	if err := json.Unmarshal(r, &data); err != nil {
		fmt.Printf("\n error in unmarshalling data in FetchDPS, error: %v", err)
		return lib.UserDPSInfo{}, err
	}

	return data.Data, nil
}

// Get the DPS of a certain discordId user.
func GetDPS(id string) (lib.UserDPSInfo, error) {
	r, err := utils.Fetcher(fmt.Sprintf("%s/dps/%s", os.Getenv("CORDY_API"), id))
	if err != nil {
		fmt.Printf("\n error in posting data in GetDPS, error: %v", err)
		return lib.UserDPSInfo{}, err
	}

	var data FetchDPSResponse
	if err := json.Unmarshal(r, &data); err != nil {
		fmt.Printf("\n error in unmarshalling data in GetDPS, error: %v", err)
		return lib.UserDPSInfo{}, err
	}

	return data.Data, nil
}

// Get the DPS of a certain discordId user. (on-demand)
func GetDPSDemand(id string) (lib.UserDPSInfo, error) {
	r, err := utils.Fetcher(fmt.Sprintf("%s/dps/demand/%s", os.Getenv("CORDY_API"), id))
	if err != nil {
		fmt.Printf("\n error in posting data in GetDPS, error: %v", err)
		return lib.UserDPSInfo{}, err
	}

	var data FetchDPSResponse
	if err := json.Unmarshal(r, &data); err != nil {
		fmt.Printf("\n error in unmarshalling data in GetDPS, error: %v", err)
		return lib.UserDPSInfo{}, err
	}

	return data.Data, nil
}
