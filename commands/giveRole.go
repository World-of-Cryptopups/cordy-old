package commands

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	e "github.com/World-of-Cryptopups/cordy/lib/errors"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

func parseRole(r string) int {
	replacer := strings.NewReplacer("<", "", ">", "", "@", "", "&", "")

	output := replacer.Replace(r)
	id, _ := strconv.Atoi(output)

	return id
}

func parseMention(m string) int {
	replacer := strings.NewReplacer("<", "", ">", "", "@", "", "!", "")

	output := replacer.Replace(m)
	id, _ := strconv.Atoi(output)

	return id
}

func (b *Bot) Giverole(c *gateway.MessageCreateEvent, role string, members ...string) (string, error) {
	if !(strings.HasPrefix(role, "<@&") && strings.HasSuffix(role, ">")) {
		return e.FailedCommand("Error parsing RoleID!", nil)
	}

	// parse role id
	roleID := discord.RoleID(parseRole(role))

	// parse each member mention
	memberIDs := []discord.UserID{}
	for _, v := range members {
		memberIDs = append(memberIDs, discord.UserID(parseMention(v)))
	}

	for _, x := range memberIDs {
		member, err := b.Ctx.Member(c.GuildID, x)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var roleExists = false
		// check for roleids
		for _, y := range member.RoleIDs {
			if y == roleID {
				roleExists = true
			}
		}

		if !roleExists {
			err = b.Ctx.AddRole(c.GuildID, x, roleID)
			log.Printf("ADDED ROLE: %d -> USER: %d || error: %v", roleID, x, err)
		} else {
			log.Printf("Role exists: %d -> User: %d", roleID, x)
		}
	}

	return fmt.Sprintf("Successfully given the role <@&%s> to them!", roleID.String()), nil
}
