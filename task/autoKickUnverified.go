package task

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
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
			// if joined within the last 3 days, wait to verify
			if time.Since(v.Joined.Time()).Hours() < 72 {
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
		// if !hasAdventureRole {
		// 	c.KickWithReason(GuildID, v.User.ID, "Unverified User")
		// }
	}

	return nil
}

func AutoKickUnverified(c *bot.Context) {
	// this is meant for auto unverified members management
	for {
		GuildID := discord.GuildID(stuff.GuildID())
		AdventureRole, _ := strconv.Atoi(os.Getenv("ADVENTURE_ROLE"))

		members, err := c.Members(GuildID)
		if err != nil {
			fmt.Println(err)
			break
		}

		handleKick(c, members, GuildID, discord.RoleID(AdventureRole))

		time.Sleep(time.Duration(1) * time.Hour)
	}
}
