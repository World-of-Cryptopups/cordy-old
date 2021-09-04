package commands

import (
	"fmt"
	"time"

	"github.com/World-of-Cryptopups/cordy/lib/db"
	e "github.com/World-of-Cryptopups/cordy/lib/errors"
	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

// Ping command.
func (b *Bot) Me(c *gateway.MessageCreateEvent) (interface{}, error) {
	b.Ctx.Typing(c.ChannelID)

	// get discordid
	_discordId := c.Author.ID.String()

	client, err := db.Client()
	if err != nil {
		return e.FailedCommand("error initializing deta db", err)
	}

	// get user (returns nil if not found)
	user, err := client.GetUser(_discordId)
	if err != nil {
		return e.FailedMessage("You are not registered! You can register by sending `>register {your-token}`.", err)
	}

	var _provider string
	if user.Type == "wax-cloud" {
		_provider = "Wax Cloud Wallet"
	} else if user.Type == "anchor" {
		_provider = "Anchor Wallet"
	}

	// >me can only be accesed by a registered user, meaning, the one who called it owns it
	// so, use the one who called it
	embed := &discord.Embed{
		Title:       fmt.Sprintf("#%d) %s", user.Rank, c.Author.Username),
		Color:       stuff.UserRoleColor(b.Ctx, c.GuildID, c.Author.ID),
		Description: "Your profile information.",
		Author: &discord.EmbedAuthor{
			Name: fmt.Sprintf("[me] %s", c.Author.Tag()),
			Icon: c.Author.AvatarURL(),
		},
		Fields: []discord.EmbedField{{
			Name:   "💳 Wallet",
			Value:  fmt.Sprintf("**%s**", user.Wallet),
			Inline: true,
		}, {
			Name:   "👥 Provider",
			Value:  _provider,
			Inline: true,
		}, {
			Name:   "🛡 Season One Pass",
			Value:  user.SeasonPasses[0].Title,
			Inline: true,
		}, {
			Name:   "🛡 Current Pass",
			Value:  fmt.Sprintf("**%s**", user.CurrentPass),
			Inline: false,
		}},
		Thumbnail: &discord.EmbedThumbnail{
			URL: c.Author.AvatarURL(),
		},
		Footer: &discord.EmbedFooter{
			Text: "© World of Cryptopups | 2021",
		},
		Timestamp: discord.Timestamp(time.Now()),
	}

	return embed, nil
}
