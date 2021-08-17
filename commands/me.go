package commands

import (
	"fmt"
	"time"

	e "github.com/World-of-Cryptopups/roleroll-new/lib/errors"
	fc "github.com/World-of-Cryptopups/roleroll-new/lib/fauna"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

// Ping command.
func (b *Bot) Me(c *gateway.MessageCreateEvent) (interface{}, error) {
	// get discordid
	_discordId := c.Author.ID.String()

	// check if user is already registered
	_registered, err := fc.CheckUser(_discordId)
	if err != nil {
		return e.FailedCommand("check if user is registered", err)
	}
	if !_registered {
		return e.FailedMessage("You are not registered! You can register by sending `>register {your-token}`.", err)
	}

	// get client
	client := fc.Client()

	// get user
	_user, err := client.Query(f.Get(f.MatchTerm(f.Index("userByDiscordId"), _discordId)))
	if err != nil {
		return e.FailedCommand("get user", err)
	}

	var user User
	if err := _user.At(f.ObjKey("data")).Get(&user); err != nil {
		return e.FailedCommand("decode data user", err)
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
		Title:       c.Author.Username,
		Description: "Your profile information.",
		Author: &discord.EmbedAuthor{
			Name: fmt.Sprintf("[me] %s", c.Author.Username),
		},
		Fields: []discord.EmbedField{{
			Name:   "💳 Wallet",
			Value:  user.DefaultWallet,
			Inline: true,
		}, {
			Name:   "👥 Provider",
			Value:  _provider,
			Inline: true,
		}, {
			Name: "\u200b", Value: "\u200b", Inline: false,
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
