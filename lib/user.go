package lib

import "github.com/World-of-Cryptopups/cordy/stuff"

//
type User struct {
	Key          string           `json:"key"` // this is needed by deta
	User         UserDiscord      `json:"user"`
	Wallet       string           `json:"wallet"`
	Type         string           `json:"type"`
	Token        string           `json:"token"`
	SeasonPasses []UserSeasonPass `json:"seasonPasses"`
}

type UserDiscord struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Tag      string `json:"tag"`
}

type UserSeasonPass struct {
	Season string           `json:"season"`
	DPS    stuff.DPSDetails `json:"dps"`
	Title  string           `json:"title"`
}

type UserCurrentSeasonPass UserSeasonPass
