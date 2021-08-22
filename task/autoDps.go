package task

import (
	"fmt"
	"strconv"
	"time"

	"github.com/World-of-Cryptopups/cordy/commands"
	fc "github.com/World-of-Cryptopups/cordy/lib/fauna"
	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

// AutoDPS is a tasks which gets the dps of the members and then resets their roles again.
func AutoDPS(c *bot.Context) {
	for {
		fmt.Println("Starting FETCHER!")

		client := fc.Client()

		x, err := client.Query(f.Map(f.Paginate(f.Documents(f.Collection("users"))), f.Lambda("x", f.Get(f.Var("x")))))
		if err != nil {
			fmt.Println(err)
		}

		var allUsers []commands.QueryUser
		x.At(f.ObjKey("data")).Get(&allUsers)

		GuildID := discord.GuildID(stuff.GuildID())

		// Loop and get again the DPS of each users registered.
		for _, v := range allUsers {
			discordID, _ := strconv.Atoi(v.Data.DiscordID)

			// check if user is in guild
			_, err := c.Member(GuildID, discord.UserID(discordID))
			if err != nil {
				// Member is not in the server, just pass him / her
				continue
			}
			fmt.Printf("[FETCHER] --> getting the data of %s", v.Data.DiscordUsername)

			if d, err := stuff.FetchDPS(stuff.UserDPSUser{
				Username: v.Data.DiscordUsername,
				Id:       v.Data.DiscordID,
				Avatar:   v.Data.AvatarURL,
			}, v.Data.DefaultWallet); err != nil {
				fmt.Printf("\n [AUTODPS] Failed Getting the DPS pof %s", v.Data.DiscordUsername)
			} else {
				totalDPS := d.DPS.Pupcards + d.DPS.Pupskins + d.DPS.Pupitems.Real

				if err := stuff.HandleUserRole(c, GuildID, discordID, totalDPS); err != nil {
					fmt.Println(err)
				}
			}

		}

		// sleep
		time.Sleep(time.Duration(5) * time.Minute)
	}
}
