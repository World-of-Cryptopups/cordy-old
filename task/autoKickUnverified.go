package task

import (
	"fmt"
	"os"
	"strconv"

	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
)

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

		for _, v := range members {
			var hasAdventureRole bool = true

			// pass if bot
			if v.User.Bot {
				continue
			}

			for _, x := range v.RoleIDs {
				if x == discord.RoleID(AdventureRole) {
					hasAdventureRole = true
					break
				}
			}

			fmt.Printf("-> %s | hasAdventureRole =>  %t\n", v.User.Tag(), hasAdventureRole)
			// if !hasAdventureRole {
			// 	c.KickWithReason(GuildID, v.User.ID, "Unverified User")
			// }
		}
	}
}
