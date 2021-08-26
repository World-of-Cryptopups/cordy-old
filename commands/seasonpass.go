package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/World-of-Cryptopups/cordy/lib"
	"github.com/World-of-Cryptopups/cordy/lib/db"
	e "github.com/World-of-Cryptopups/cordy/lib/errors"
	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/World-of-Cryptopups/cordy/utils"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

var CurrentSeasons = []string{"one"}

func (b *Bot) Seasonpass(c *gateway.MessageCreateEvent, args bot.RawArguments) (interface{}, error) {
	b.Ctx.Typing(c.ChannelID)

	// default message without args
	season := string(args)

	// get discordid
	_discordId := c.Author.ID.String()

	// check if user is already registered
	// if _, err := fc.IsUserRegistered(_discordId); err != nil {
	// 	return "", err
	// }

	// _query, err := client.Query(f.Let().Bind("user", f.Get(f.MatchTerm(f.Index("userByDiscordId"), _discordId))).In(
	// 	f.Map(f.Paginate(f.MatchTerm(f.Index("passByUser"), f.Select("ref", f.Var("user"))), f.Size(1)),
	// 		f.Lambda("dps", f.Get(f.Var("dps"))))))
	// if err != nil {
	// 	return e.FailedCommand("get user seasonpass dps", err)
	// }

	// var _dps []QueryUserSeasonPass
	// if err := _query.At(f.ObjKey("data")).Get(&_dps); err != nil {
	// 	return e.FailedCommand("decode dps response", err)
	// }

	client, err := db.Client()
	if err != nil {
		return e.FailedCommand("failed initializing Deta Base", err)
	}

	// check if user already exists
	user, err := client.GetUser(_discordId)
	if err != nil {
		return e.FailedMessage("You are not registered! You can register by sending `>register {your-token}`.", err)
	}

	// season not specified
	if season == "" {
		embed := &discord.Embed{
			Title:       c.Author.Username,
			Color:       stuff.UserRoleColor(b.Ctx, c.GuildID, c.Author.ID),
			Description: "Your current season pass.",
			Author: &discord.EmbedAuthor{
				Name: fmt.Sprintf("[me] %s", c.Author.Tag()),
				Icon: c.Author.AvatarURL(),
			},
			Fields: []discord.EmbedField{{
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

	// args is present but does not exist on season
	if !utils.StringContains(CurrentSeasons, season) {
		return e.FailedMessage("We don't have that season yet, if there is a problem, you can message a mod for more info.", nil)
	}

	var data lib.UserSeasonPass
	for _, v := range user.SeasonPasses {
		if strings.EqualFold(v.Season, season) {
			data = v
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
