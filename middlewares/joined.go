package middlewares

import (
	"errors"

	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/bot/extras/infer"
	"github.com/diamondburned/arikawa/v2/discord"
)

func DisallowNotJoined(ctx *bot.Context) func(interface{}) error {
	return func(ev interface{}) error {
		// Try and infer the GuildID.
		guildID := infer.GuildID(ev)
		if !guildID.IsValid() {
			return errors.New("not joined")
		}

		if guildID != discord.GuildID(stuff.GuildID()) {
			return bot.Break
		}

		return nil
	}
}
