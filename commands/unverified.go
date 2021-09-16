package commands

import (
	"fmt"
	"os"
	"strconv"
	"time"

	e "github.com/World-of-Cryptopups/cordy/lib/errors"
	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

func handleKick(c *bot.Context, members []discord.Member, guildID discord.GuildID, AdventureRole discord.RoleID) error {
	fmt.Printf("total members: %d\n", len(members))
	for i, v := range members {
		var hasAdventureRole bool = false

		// pass if bot
		if v.User.Bot {
			continue
		}

		for _, x := range v.RoleIDs {
			if AdventureRole == x {
				hasAdventureRole = true
				break
			}
		}

		if !hasAdventureRole {
			// if joined within the last 1 day, wait to verify
			if time.Since(v.Joined.Time()).Hours() < 24 {
				continue
			}

			fmt.Printf("-> %s | hasAdventureRole =>  %t\n", v.User.Tag(), hasAdventureRole)
		}

		if i == len(members)-1 {
			if i == 999 {
				m, err := c.MembersAfter(guildID, v.User.ID, 1000)
				if err != nil {
					fmt.Println(err)
					break
				}

				return handleKick(c, m, guildID, AdventureRole)
			}
		}
		if !hasAdventureRole {
			c.KickWithReason(guildID, v.User.ID, "Unverified User")
		}
	}

	return nil
}

// KickUnverified is a special command to kick members that hasn't verified yet.
func (b *Bot) KickUnverified(c *gateway.MessageCreateEvent) (interface{}, error) {
	b.Ctx.Typing(c.ChannelID)

	GuildID := discord.GuildID(stuff.GuildID())
	AdventureRole, _ := strconv.Atoi(os.Getenv("ADVENTURE_ROLE"))

	members, err := b.Ctx.Members(GuildID)
	if err != nil {
		return e.FailedCommand("failed to get all members", err)
	}

	handleKick(b.Ctx, members, GuildID, discord.RoleID(AdventureRole))

	return "Successfully removed all unverified members!", nil
}
