package lib

type UserDPSInfo struct {
	Wallet string      `json:"wallet"`
	Key    string      `json:"key,omitempty"`
	User   UserDPSUser `json:"user"`
	DPS    DPSDetails  `json:"dps"`
}

type UserDPSUser UserDiscord
