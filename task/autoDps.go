package task

import (
	"fmt"
	"time"

	"github.com/World-of-Cryptopups/cordy/commands"
	fc "github.com/World-of-Cryptopups/cordy/lib/fauna"
	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/diamondburned/arikawa/v2/bot"
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

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

		// Loop and get again the DPS of each users registered.
		for _, v := range allUsers {
			stuff.FetchDPS(stuff.UserDPSUser{
				Username: v.Data.DiscordUsername,
				Id:       v.Data.DiscordID,
				Avatar:   v.Data.AvatarURL,
			}, v.Data.DefaultWallet)
		}

		// sleep
		time.Sleep(time.Duration(1) * time.Minute)
	}
}
