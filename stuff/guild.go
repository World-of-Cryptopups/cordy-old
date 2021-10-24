package stuff

import (
	"fmt"
	"os"
	"strconv"

	"github.com/World-of-Cryptopups/cordy/utils"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
)

// GuildID is the servers' id, (CHANGE THIS IN PRODUCTION)
func GuildID() int {
	var g, _ = strconv.Atoi(os.Getenv("DEFAULT_GUILDID"))
	return g
}

func UserRoleColor(b *bot.Context, guildID discord.GuildID, userID discord.UserID) discord.Color {
	// get role info
	var _embedColor string

	// get member
	member, err := b.Client.Member(guildID, userID)
	if err != nil {
		fmt.Println(err)
		return discord.Color(0)
	}

	if role := GetHighestRole(member); role.Title == "" {
		// default to grey color
		_embedColor = Colors["grey"]
	} else {
		_embedColor = Colors[role.Color]
	}

	return discord.Color(utils.ParseColor(_embedColor))
}
