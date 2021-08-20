// Package commands is where all bot commands go and being managed
package commands

import (
	"github.com/World-of-Cryptopups/roleroll-new/stuff"
	"github.com/diamondburned/arikawa/v2/bot"
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

type Bot struct {
	Ctx *bot.Context
}

type User struct {
	DiscordID       string             `fauna:"discordId,omitempty"`
	DiscordUsername string             `fauna:"discordUsername,omitempty"`
	AvatarURL       string             `fauna:"avatarUrl,omitempty"`
	Wallets         []string           `fauna:"wallets,omitempty"`
	DefaultWallet   string             `fauna:"defaultWallet,omitempty"`
	Type            string             `fauna:"type,omitempty"`
	Token           string             `fauna:"token,omitempty"`
	SeasonPasses    []UserSeasonPasses `fauna:"seasonPasses"`
}

type QueryUser struct {
	Ref  f.RefV `fauna:"ref"`
	Data User   `fauna:"data"`
}

type UserSeasonPasses struct {
	Season string `fauna:"season"`
	Title  string `fauna:"title"`
}

type UserSeasonPass struct {
	User   f.RefV           `fauna:"user"`
	Season string           `fauna:"season,omitempty"`
	DPS    stuff.DPSDetails `fauna:"dps,omitempty"`
	Title  string           `fauna:"title,omitempty"`
}

type QueryUserSeasonPass struct {
	Ref  f.RefV         `fauna:"ref"`
	Data UserSeasonPass `fauna:"data"`
}
