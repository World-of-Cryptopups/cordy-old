// Package commands is where all bot commands go and being managed
package commands

import (
	"time"

	"github.com/World-of-Cryptopups/cordy/middlewares"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

type Bot struct {
	Ctx *bot.Context
}

func (b *Bot) Setup(sub *bot.Subcommand) {
	// do not allow dm messages
	sub.AddMiddleware("*", middlewares.DisallowNotJoined(b.Ctx))
}

// Help returns the help message for the bot.
func (b *Bot) Help(c *gateway.MessageCreateEvent) (interface{}, error) {
	me, _ := b.Ctx.Me()

	// c.Type == discord.MessageType(discord.DirectMessage)

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
