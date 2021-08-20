package stuff

import (
	"encoding/json"
	"os"

	"github.com/World-of-Cryptopups/cordy/utils"
)

type UserDPSInfo struct {
	Wallet string      `json:"wallet"`
	Key    string      `json:"key,omitempty"`
	User   UserDPSUser `json:"user"`
	DPS    DPSDetails  `json:"dps"`
}

type UserDPSUser struct {
	Avatar   string `json:"avatar"`
	Id       string `json:"id"`
	Username string `json:"username"`
}

// FetchDPS is a fetcher to call the endpoint and save to db.
func FetchDPS(user UserDPSUser, wallet string) (UserDPSInfo, error) {
	r, err := utils.PostFetcher(user, os.Getenv("DPS_FETCH")+wallet)
	if err != nil {
		return UserDPSInfo{}, err
	}

	var data UserDPSInfo
	if err := json.Unmarshal(r, &data); err != nil {
		return UserDPSInfo{}, err
	}

	return data, nil
}

// Get the DPS of a certain discordId user.
func GetDPS(id string) (UserDPSInfo, error) {
	r, err := utils.Fetcher(os.Getenv("DPS_GET") + id)
	if err != nil {
		return UserDPSInfo{}, err
	}

	var data UserDPSInfo
	if err := json.Unmarshal(r, &data); err != nil {
		return UserDPSInfo{}, err
	}

	return data, nil
}
