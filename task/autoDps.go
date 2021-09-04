package task

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/World-of-Cryptopups/cordy/lib"
	"github.com/World-of-Cryptopups/cordy/lib/db"
	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/deta/deta-go/service/base"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
)

type UserRankDPS struct {
	UserID     string // discord id
	UserAvatar string // sometimes, the user auto-updates their avatar, so.. update 'em also
	Wallet     string
	TotalDPS   int
}

// AutoDPS is a tasks which gets the dps of the members and then resets their roles again.
func AutoDPS(c *bot.Context) {
	for {
		fmt.Println("Starting FETCHER!")

		client, _ := db.Client()

		users, err := client.GetAllUsers()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// users ranking
		usersRanking := []UserRankDPS{}

		fmt.Printf("\nTOTAL USERS: %d", len(users))

		GuildID := discord.GuildID(stuff.GuildID())

		// Loop and get again the DPS of each users registered.
		for _, v := range users {
			discordId, _ := strconv.Atoi(v.User.ID)

			// check if user is in guild
			member, err := c.Member(GuildID, discord.UserID(discordId))
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

				// include to dps ranking slice
				usersRanking = append(usersRanking, UserRankDPS{
					UserID:     v.User.ID,
					UserAvatar: member.User.Avatar,
					Wallet:     v.Wallet,
					TotalDPS:   totalDPS,
				})
			}
			// sleep for 2 seconds
			time.Sleep(time.Duration(2) * time.Second)
		}

		// sort `usersRanking`
		sort.SliceStable(usersRanking, func(i, j int) bool {
			return usersRanking[i].TotalDPS > usersRanking[j].TotalDPS
		})

		// loop sorted slice to update user rankings & others
		for index, v := range usersRanking {
			// get the current pass
			pass, err := stuff.GetCurrentPass(v.Wallet)
			if err != nil {
				fmt.Println(err)
			}

			if err = client.DB.Update(v.UserID, base.Updates{
				"currentPass": pass.Pass, // just update the pass, xD
				"user.avatar": v.UserAvatar,
				"rank":        index + 1,
			}); err != nil {
				fmt.Println("failed to update user info")
			}
		}

		// sleep
		time.Sleep(time.Duration(5) * time.Minute)
	}
}
