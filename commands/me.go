package commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/World-of-Cryptopups/cordy/lib/db"
	e "github.com/World-of-Cryptopups/cordy/lib/errors"
	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/diamondburned/arikawa/v2/bot"
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
		if who == "me" {
			return e.RegisterErr()
		}

		return "User is not yet registered!", nil
	}

	var _provider string
	if user.Type == "wax-cloud" {
		_provider = "Wax Cloud Wallet"
	} else if user.Type == "anchor" {
		_provider = "Anchor Wallet"
	}

	userID, _ := strconv.Atoi(user.User.ID)

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
func (b *Bot) Info(c *gateway.MessageCreateEvent, args bot.RawArguments) (interface{}, error) {
	b.Ctx.Typing(c.ChannelID)

	uargs := string(args)
	if uargs == "" {
		return "No users to get info!", nil
	}

	usermentions := strings.Split(uargs, " ")

	fmt.Println(len(usermentions))
	fmt.Println(usermentions)

	if len(usermentions) == 1 {
		_discordId := strconv.Itoa(parseMention(usermentions[0]))

		if _discordId == "0" || _discordId == "" {
			return e.FailedMessage("Unknown user!", nil)
		}

		return b.getUserInfo(_discordId, c.GuildID, "user")
	}

	for _, v := range usermentions {
		_discordId := strconv.Itoa(parseMention(v))

		if _discordId == "0" || _discordId == "" {
			continue
			// return e.FailedMessage("Unknown user!", nil)
		}

		em, err := b.getUserInfo(_discordId, c.GuildID, "user")
		if err != nil {
			continue
		}

		embed, ok := em.(*discord.Embed)
		if !ok {
			continue
		}

		b.Ctx.SendMessage(c.ChannelID, "", embed)
	}

	return nil, nil
}
