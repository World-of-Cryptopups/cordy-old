package main

import (
	"log"
	"os"

	"github.com/World-of-Cryptopups/roleroll-new/commands"
	_ "github.com/joho/godotenv/autoload"

	"github.com/diamondburned/arikawa/v2/bot"
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

		return nil
	})

}
