package commands

import (
	"fmt"
	"strings"

	"github.com/World-of-Cryptopups/cordy/lib"
	"github.com/World-of-Cryptopups/cordy/lib/db"
	e "github.com/World-of-Cryptopups/cordy/lib/errors"
	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

func (b *Bot) List(c *gateway.MessageCreateEvent, args bot.RawArguments) (interface{}, error) {
	b.Ctx.Typing(c.ChannelID)

	uargs := string(args)

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

	var lists = map[string][]*lib.UserDPSInfo{}
	for _, v := range datas {
		totalDps := v.DPS.Pupcards + v.DPS.Pupskins + v.DPS.Pupitems.Real

		_role := stuff.GetDPSRoleInfo(totalDps)
		lists[_role.Title] = append(lists[_role.Title], v)

	}

	dm, err := b.Ctx.CreatePrivateChannel(c.Author.ID)
	if err != nil {
		return e.FailedMessage("Failed trying to send list through private / dm message.", err)
	}

	for i, k := range lists {
		_members := []string{}
		for _, j := range k {

			if uargs == "wallets" {
				_members = append(_members, j.Wallet)
				continue
			}

			_members = append(_members, j.User.Tag)
		}

		role := i
		if role == "" {
			role = "Adventure Pups"
		}

		embed := &discord.Embed{
			Title:       fmt.Sprintf("List of **%s** Members", role),
			Description: strings.Join(_members, "\n"),
			Author: &discord.EmbedAuthor{
				Name: me.Username,
				Icon: me.AvatarURL(),
			},
			Thumbnail: &discord.EmbedThumbnail{
				URL: me.AvatarURL(),
			},
		}

		b.Ctx.SendMessage(dm.ID, "", embed)

	}

	return "I have sent you a private message with the list.", nil
}
