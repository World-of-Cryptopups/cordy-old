// Package commands is where all bot commands go and being managed
package commands

import (
	"time"

	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
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

// Help returns the help message for the bot.
func (b *Bot) Help(*gateway.MessageCreateEvent) (interface{}, error) {
	me, _ := b.Ctx.Me()

	embed := &discord.Embed{
		Author: &discord.EmbedAuthor{
			Name: me.Username,
			Icon: me.AvatarURL(),
		},
		Title:       "Usage | Cordy Commands",
		Description: "The following are my commands, if you don't know what to do, please contact an admin or mod for more info.",
		Fields: []discord.EmbedField{
			{
				Name:  "**`>register`**",
				Value: "Register your account to the Bot. Token is provided at https://www.worldofcryptopups.cf/my-collections. \nUsage: `>register [token]`",
			},
			{
				Name:  "**`>help`**",
				Value: "Show this help message.",
			},
			{
				Name:  "**`>dps`**",
				Value: "*[only-registered]* Get your current DPS info. The data could be delayed and be different from what our website shows.",
			},
			{
				Name:  "**`>me`**",
				Value: "*[only-registered]* Show information about your account.",
			},
			{
				Name:  "**`>seasonpass`**",
				Value: "*[only-registered]* Get your DPS on a specific season. \nExample: `>seasonpass one`",
			},
		},
		Thumbnail: &discord.EmbedThumbnail{
			URL: me.AvatarURL(),
		},
		Footer: &discord.EmbedFooter{
			Text: "© World of Cryptopups | 2021",
		},
		Timestamp: discord.Timestamp(time.Now()),
		URL:       "https://www.worldofcryptopups.cf/",
	}

	return embed, nil
}
