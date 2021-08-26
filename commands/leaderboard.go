package commands

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/World-of-Cryptopups/cordy/lib/db"
	e "github.com/World-of-Cryptopups/cordy/lib/errors"

	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

func (b *Bot) Leaderboard(c *gateway.MessageCreateEvent) (interface{}, error) {
	b.Ctx.Typing(c.ChannelID)

	me, _ := b.Ctx.Me()

	client, err := db.Client()
	if err != nil {
		return e.FailedCommand("error initializing deta db", err)
	}

	// get all dps infos
	datas, err := client.GetAllDPS()
	if err != nil {
		return e.FailedCommand("error in getting all the dps data", err)
	}

	// sort
	sort.SliceStable(datas, func(i, j int) bool {
		_iDPS := datas[i].DPS.Pupcards + datas[i].DPS.Pupskins + datas[i].DPS.Pupitems.Real
		_jDPS := datas[j].DPS.Pupcards + datas[j].DPS.Pupskins + datas[j].DPS.Pupitems.Real

		return _iDPS > _jDPS
	})

	// get top10
	var top []discord.EmbedField

	for i, v := range datas[:10] {
		_total := v.DPS.Pupcards + v.DPS.Pupskins + v.DPS.Pupitems.Real

		top = append(top, discord.EmbedField{
			Name:  fmt.Sprintf("%d) %s", i+1, v.User.Tag),
			Value: fmt.Sprintf("🛡 **%s** DPS", strconv.Itoa(_total)),
		})
	}

	embed := &discord.Embed{
		Author: &discord.EmbedAuthor{
			Name: me.Username,
			Icon: me.AvatarURL(),
		},
		Title:       "Members Leaderboard",
		Description: "Here are the top members, *(only registered members are added)*",
		Fields:      top,
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
