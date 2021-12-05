package commands

import (
	"fmt"
	"strconv"

	"github.com/World-of-Cryptopups/cordy/lib/db"
	e "github.com/World-of-Cryptopups/cordy/lib/errors"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

func (b *Bot) removeUser(discordId string) (interface{}, error) {
	db, _ := db.Client()

	if exists, err := db.UserExists(discordId); err != nil {
		if !exists {
			// if it does not exist, do not continue
			return e.FailedMessage("User does not exist or is not logged in.", err)
		}
	} else {
		return e.FailedMessage("There was a problem checking the user if it is logged in or not.", err)
	}

	if err := db.DB.Delete(discordId); err != nil {
		return e.FailedCommand("remove user", err)
	}
	if err := db.DPSDB.Delete(discordId); err != nil {
		return e.FailedCommand("remove user dps", err)
	}

	return fmt.Sprintf("Account of <@%s> has been resetted successfully.", discordId), nil

}

// Admin only command for resetting the wallet of the user.
// This command removes the logged in wallet auth of the discord user id.
func (b *Bot) ResetAccount(c *gateway.MessageCreateEvent, args bot.RawArguments) (interface{}, error) {
	b.Ctx.Typing(c.ChannelID)

	uargs := string(args)
	if uargs == "" {
		return "This command will reset the account of the user mentioned.", nil
	}

	_discordId := strconv.Itoa(parseMention(uargs))
	userID, _ := strconv.Atoi(_discordId)

	_, err := b.Ctx.Member(c.GuildID, discord.UserID(userID))
	if err != nil {
		return e.FailedMessage("Failed to get user.", err)
	}
	userpermissions, err := b.Ctx.Permissions(c.ChannelID, discord.UserID(userID))
	if err != nil {
		return e.FailedMessage("Failed to get user's permissions.", err)
	}

	// if user (mentioned) has admin permission, do not reset
	if userpermissions.Has(discord.PermissionAdministrator) {
		return e.FailedMessage("I cannot reset the account of ana admin.", err)
	}

	if _discordId == "0" || _discordId == "" {
		return e.FailedMessage("Unknown user!", nil)
	}

	return b.removeUser(_discordId)
}
