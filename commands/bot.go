// Package commands is where all bot commands go and being managed
package commands

import (
	"github.com/diamondburned/arikawa/v2/bot"
)

type Bot struct {
	Ctx *bot.Context
}

type User struct {
	DiscordID     string   `fauna:"discordId,omitempty"`
	AvatarURL     string   `fauna:"avatarUrl,omitempty"`
	Wallets       []string `fauna:"wallets,omitempty"`
	DefaultWallet string   `fauna:"defaultWallet,omitempty"`
	Type          string   `fauna:"type,omitempty"`
	Token         string   `fauna:"token,omitempty"`
}
