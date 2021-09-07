package main

import (
	"log"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload" // this should be first

	"github.com/World-of-Cryptopups/cordy/commands"
	"github.com/World-of-Cryptopups/cordy/task"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

func main() {
	var token = os.Getenv("TOKEN")

	if token == "" {
		log.Fatalln("Missing TOKEN!")
	}

	commands := &commands.Bot{}

	bot.Run(token, commands, func(ctx *bot.Context) error {
		ctx.HasPrefix = bot.NewPrefix(">")
		ctx.EditableCommands = true

		// DO NOT SEND `unknown command`
		ctx.SilentUnknown.Command = true
		ctx.SilentUnknown.Subcommand = true

		ctx.Gateway.AddIntents(gateway.IntentDirectMessages)
		ctx.Gateway.AddIntents(gateway.IntentGuildMessages)
		ctx.Gateway.AddIntents(gateway.IntentGuilds)

		// change status
		ctx.Gateway.Identifier.IdentifyData = gateway.IdentifyData{
			Token: ctx.Token,
			Presence: &gateway.UpdateStatusData{
				Activities: []discord.Activity{
					{
						Name: "World of Cryptopups",
						Type: discord.WatchingActivity,
					},
				},
			},
		}

		// DO NOT RUN THE FETCHER ON DEVELOPMENT MODE
		if dev, _ := strconv.ParseBool(os.Getenv("DEV_MODE")); dev {
			// run task (disable for now)
			go task.AutoDPS(ctx)
		}

		return nil
	})

}
