package task

import (
	"fmt"
	"strconv"
	"time"

	"github.com/World-of-Cryptopups/cordy/lib"
	"github.com/World-of-Cryptopups/cordy/lib/db"
	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
)

// AutoDPS is a tasks which gets the dps of the members and then resets their roles again.
func AutoDPS(c *bot.Context) {
	for {
		fmt.Println("Starting FETCHER!")

		client, _ := db.Client()

		users, err := client.GetAllUsers()
		if err != nil {
			fmt.Println(err)
		}

		GuildID := discord.GuildID(stuff.GuildID())

		// Loop and get again the DPS of each users registered.
		for _, v := range users {
			discordId, _ := strconv.Atoi(v.User.ID)

			// check if user is in guild
			_, err := c.Member(GuildID, discord.UserID(discordId))
			if err != nil {
				// Member is not in the server, just pass him / her
				continue
			}
			fmt.Printf("\n[FETCHER] --> getting the data of %s", v.User.Username)

			if d, err := stuff.FetchDPS(lib.UserDPSUser{
				Username: v.User.Username,
				Tag:      v.User.Tag,
				ID:       v.User.ID,
				Avatar:   v.User.Avatar,
			}, v.Wallet); err != nil {
				fmt.Println(err)
				fmt.Printf("\n [AUTODPS] Failed Getting the DPS pof %s", v.User.Username)
			} else {
				totalDPS := d.DPS.Pupcards + d.DPS.Pupskins + d.DPS.Pupitems.Real

				if err := stuff.HandleUserRole(c, GuildID, discordId, totalDPS); err != nil {
					fmt.Println(err)
				}
			}

		}

		// sleep
		time.Sleep(time.Duration(5) * time.Minute)
	}
}
