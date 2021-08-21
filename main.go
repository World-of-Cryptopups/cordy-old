package main

import (
	"log"
	"os"
	"strconv"

	"github.com/World-of-Cryptopups/cordy/commands"
	"github.com/World-of-Cryptopups/cordy/task"
	_ "github.com/joho/godotenv/autoload"

	"github.com/diamondburned/arikawa/v2/bot"
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

		// DO NOT RUN THE FETCHER ON DEVELOPMENT MODE
		if dev, _ := strconv.ParseBool(os.Getenv("DEV_MODE")); !dev {
			// run task (disable for now)
			go task.AutoDPS(ctx)
		}

		return nil
	})

}
