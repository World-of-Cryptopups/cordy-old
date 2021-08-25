package lib

type SeasonPassVerify struct {
	Pass   string `json:"pass"`
	Wallet string `json:"wallet"`
	Season string `json:"season"`
}

type SeasonPass struct {
	Wallet string     `json:"wallet"`
	Season string     `json:"season"`
	DPS    DPSDetails `json:"dps"`
}
type DPSDetails struct {
	Pupskins int `json:"pupskins" fauna:"pupskins"`
	Pupcards int `json:"pupcards" fauna:"pupcards"`
	Pupitems struct {
		Raw  int `json:"raw" fauna:"raw"`
		Real int `json:"real" fauna:"real"`
	} `json:"pupitems" fauna:"pupitems"`
}
