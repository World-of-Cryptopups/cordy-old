package commands

import (
	"errors"
	"fmt"

	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/enescakir/emoji"
)

func (b *Bot) Register(c *gateway.MessageCreateEvent, f bot.RawArguments) (string, error) {
	if f == "" {
		return "", errors.New("no TOKEN provided")
	}

	return fmt.Sprintf("%v Successfully authenticated <@!%s>!", emoji.CheckBoxWithCheck, c.Author.ID), nil
}
