package commands

import (
	"fmt"
	"time"

	"github.com/World-of-Cryptopups/cordy/lib/db"
	e "github.com/World-of-Cryptopups/cordy/lib/errors"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/enescakir/emoji"
)

func (b *Bot) Stats(c *gateway.MessageCreateEvent) (interface{}, error) {
	b.Ctx.Typing(c.ChannelID)

	me, _ := b.Ctx.Me()
	client, _ := db.Client()

	users, err := client.GetAllUsers()
	if err != nil {
		return e.FailedCommand("error getting all users", err)
	}

	embed := &discord.Embed{
		Author: &discord.EmbedAuthor{
			Name: me.Username,
			Icon: me.AvatarURL(),
		},
		Title:       "World of Cryptopups Stats",
		Description: "Cordy bot current statistics.",
		Fields: []discord.EmbedField{{
			Name:  fmt.Sprintf("%v Total Registered Users", emoji.BustsInSilhouette),
			Value: fmt.Sprintf("%d", len(users)),
		}},
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
