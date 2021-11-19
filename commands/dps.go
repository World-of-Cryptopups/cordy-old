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

func (b *Bot) Dps(c *gateway.MessageCreateEvent) (interface{}, error) {
	b.Ctx.Typing(c.ChannelID)

	// get discordid
	_discordId := c.Author.ID.String()

	client, err := db.Client()
	if err != nil {
		return e.FailedCommand("error initializing deta db", err)
	}

	// get user
	userExists, err := client.UserExists(_discordId)
	if err != nil {
		return e.FailedCommand("error getting the user from db", err)
	}
	if !userExists {
		return e.RegisterErr()
	}

	// get dps info
	data, err := stuff.GetDPSDemand(c.Author.ID.String())
	if err != nil {
		return e.FailedCommand("error getting dps", err)
	}

	totalDps := data.DPS.Pupcards + data.DPS.Pupskins + data.DPS.Pupitems.Real

	embed := &discord.Embed{
		Author: &discord.EmbedAuthor{
			Name: c.Author.Tag(),
			Icon: c.Author.AvatarURL(),
		},
		Color:       stuff.UserRoleColor(b.Ctx, c.GuildID, c.Author.ID),
		Title:       "Current DPS Stats",
		Description: "*(If data is wrong, please report to a mod or admin.)*",
		// Description: fmt.Sprintf("Your total **DPS** accumulated for Season %s", strings.Title(data.Season)),
		Thumbnail: &discord.EmbedThumbnail{
			URL: c.Author.AvatarURL(),
		},
		Fields: []discord.EmbedField{
			{
				Name:   "🎴 Puppy Cards",
				Value:  fmt.Sprint(data.DPS.Pupcards),
				Inline: true,
			},
			{
				Name:   "🃏 Pup Skins",
				Value:  fmt.Sprint(data.DPS.Pupskins),
				Inline: true,
			},
			{
				Name:   "⚔️ Pup Items (Raw)",
				Value:  fmt.Sprint(data.DPS.Pupitems.Raw),
				Inline: true,
			},
			{
				Name:   "⚔️ Pup Items (Real)",
				Value:  fmt.Sprint(data.DPS.Pupitems.Real),
				Inline: true,
			},
			{
				Name:  "\u200b",
				Value: "\u200b",
			},
			{
				Name:  "🛡 Total DPS",
				Value: fmt.Sprintf("**%d**", totalDps),
			},
		},
		Footer: &discord.EmbedFooter{
			Text: "© World of Cryptopups | 2021",
		},
		Timestamp: discord.Timestamp(time.Now()),
	}

	return embed, nil
}
