package commands

import (
	"fmt"
	"strings"
	"time"

	e "github.com/World-of-Cryptopups/cordy/lib/errors"
	fc "github.com/World-of-Cryptopups/cordy/lib/fauna"
	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/World-of-Cryptopups/cordy/utils"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

var CurrentSeasons = []string{"one"}

func (b *Bot) Seasonpass(c *gateway.MessageCreateEvent, args bot.RawArguments) (interface{}, error) {
	b.Ctx.Typing(c.ChannelID)

	// default message without args
	season := string(args)
	if season == "" {
		return "**Season Pass DPS** is the total DPS accumulated after a season, depends on time.", nil
	}

	// args is present but does not exist on season
	if !utils.StringContains(CurrentSeasons, season) {
		return e.FailedMessage("We don't have that season yet, if there is a problem, you can message a mod for more info.", nil)
	}

	// get discordid
	_discordId := c.Author.ID.String()

	// check if user is already registered
	if _, err := fc.IsUserRegistered(_discordId); err != nil {
		return "", err
	}

	client := fc.Client()

	_query, err := client.Query(f.Let().Bind("user", f.Get(f.MatchTerm(f.Index("userByDiscordId"), _discordId))).In(
		f.Map(f.Paginate(f.MatchTerm(f.Index("passByUser"), f.Select("ref", f.Var("user"))), f.Size(1)),
			f.Lambda("dps", f.Get(f.Var("dps"))))))
	if err != nil {
		return e.FailedCommand("get user seasonpass dps", err)
	}

	var _dps []QueryUserSeasonPass
	if err := _query.At(f.ObjKey("data")).Get(&_dps); err != nil {
		return e.FailedCommand("decode dps response", err)
	}

	var data UserSeasonPass
	for _, v := range _dps {
		if v.Data.Season == season {
			data = v.Data
		}
	}

	totalDps := data.DPS.Pupcards + data.DPS.Pupskins + data.DPS.Pupitems.Real

	// 	p := message.NewPrinter(message.MatchLanguage("en"))
	embed := &discord.Embed{
		Author: &discord.EmbedAuthor{
			Name: c.Author.Tag(),
			Icon: c.Author.AvatarURL(),
		},
		Color: stuff.UserRoleColor(b.Ctx, c.GuildID, c.Author.ID),
		Title: fmt.Sprintf("Season %s Pass", strings.Title(data.Season)),
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
