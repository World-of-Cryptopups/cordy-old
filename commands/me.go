package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/World-of-Cryptopups/cordy/lib/db"
	e "github.com/World-of-Cryptopups/cordy/lib/errors"
	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

func (b *Bot) getUserInfo(userid string, guildID discord.GuildID, who string) (interface{}, error) {
	client, err := db.Client()
	if err != nil {
		return e.FailedCommand("error initializing deta db", err)
	}

	// get user (returns nil if not found)
	user, err := client.GetUser(userid)
	if err != nil {
		return e.RegisterErr()
	}

	var _provider string
	if user.Type == "wax-cloud" {
		_provider = "Wax Cloud Wallet"
	} else if user.Type == "anchor" {
		_provider = "Anchor Wallet"
	}

	userID, _ := strconv.Atoi(user.User.ID)

	fmt.Println(user.User)

	// >me can only be accesed by a registered user, meaning, the one who called it owns it
	// so, use the one who called it
	embed := &discord.Embed{
		Color:       stuff.UserRoleColor(b.Ctx, guildID, discord.UserID(userID)),
		Description: "Your profile information.",
		Author: &discord.EmbedAuthor{
			Name: fmt.Sprintf("[%s] %s", who, user.User.Tag),
			Icon: user.User.Avatar,
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
			Name:   "🛡 Current Pass",
			Value:  fmt.Sprintf("**%s**", user.CurrentPass),
			Inline: true,
		}},
		Thumbnail: &discord.EmbedThumbnail{
			URL: user.User.Avatar,
		},
		Footer: &discord.EmbedFooter{
			Text: "© World of Cryptopups | 2021",
		},
		Timestamp: discord.Timestamp(time.Now()),
	}

	if user.Rank != 0 {
		// has dps
		embed.Title = fmt.Sprintf("#%d", user.Rank)
	} else if user.Rank == -1 {
		// no dps
		embed.Title = "(unranked)"
	} else {
		// waiting for calculation
		embed.Title = "(unranked - waiting)"
	}

	fmt.Println(embed)

	return embed, nil
}

// Ping command.
func (b *Bot) Me(c *gateway.MessageCreateEvent) (interface{}, error) {
	b.Ctx.Typing(c.ChannelID)

	// get discordid
	_discordId := c.Author.ID.String()

	return b.getUserInfo(_discordId, c.GuildID, "me")
}

// Info gets the user info a specific member.
func (b *Bot) Info(c *gateway.MessageCreateEvent, userMention string) (interface{}, error) {
	b.Ctx.Typing(c.ChannelID)

	_discordId := strconv.Itoa(parseMention(userMention))

	if _discordId == "0" || _discordId == "" {
		return e.FailedMessage("Unknown user!", nil)
	}

	return b.getUserInfo(_discordId, c.GuildID, "user")
}
