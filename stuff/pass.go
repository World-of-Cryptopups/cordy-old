package stuff

type SeasonPassVerify struct {
	Pass   string `json:"pass"`
	Wallet string `json:"wallet"`
	Season string `json:"season"`
}

type SeasonPass struct {
	Wallet string        `json:"wallet"`
	Season string        `json:"season"`
	DPS    SeasonPassDPS `json:"dps"`
}
type SeasonPassDPS struct {
	Pupskins int `json:"pupskins" fauna:"pupskins"`
	Pupcards int `json:"pupcards" fauna:"pupcards"`
	Pupitems struct {
		Raw  int `json:"raw" fauna:"raw"`
		Real int `json:"real" fauna:"real"`
	} `json:"pupitems" fauna:"pupitems"`
}
