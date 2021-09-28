package commands

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	e "github.com/World-of-Cryptopups/cordy/lib/errors"
	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
)

func (b *Bot) handleKick(members []discord.Member, guildID discord.GuildID) error {
	fmt.Printf("total members: %d\n", len(members))

	uMembers := b.getUnverifiedMembers(guildID, members)
	for _, v := range uMembers {
		fmt.Println("Member: ", v.User.Username)

		b.Ctx.KickWithReason(guildID, v.User.ID, "Unverified User")
	}

	return nil
}

// get a list of unverified members
// members/users that don't have `Adventure Pup` as a role
func (b *Bot) getUnverifiedMembers(guildID discord.GuildID, members []discord.Member) []discord.Member {
	AdventureRole, _ := strconv.Atoi(os.Getenv("ADVENTURE_ROLE"))

	return b.filterMembers([]discord.Member{}, guildID, members, discord.RoleID(AdventureRole))

}

// checks if the adventurerole is in roleids
func hasAdventureRole(roleids []discord.RoleID, adventRole discord.RoleID) bool {
	for _, v := range roleids {
		if v == adventRole {
			return true
		}
	}

	return false
}

func (b *Bot) filterMembers(uMembers []discord.Member, guildID discord.GuildID, members []discord.Member, adventRole discord.RoleID) []discord.Member {
	for i, v := range members {
		// if user is bot, pass
		if v.User.Bot {
			continue
		}

		if !hasAdventureRole(v.RoleIDs, adventRole) {
			// if joined within the last 1 day, wait to verify
			// prevents kicking of users that are newly joined
			if time.Since(v.Joined.Time()).Hours() < 24 {
				continue
			}

			uMembers = append(uMembers, v)
		}

		if i == len(members)-1 {
			if i == 999 {
				m, err := b.Ctx.MembersAfter(guildID, v.User.ID, 1000)
				if err != nil {
					fmt.Println(err)
					break
				}

				return b.filterMembers(uMembers, guildID, m, adventRole)
			}
		}

	}

	return uMembers

}

func JoinMemberMentions(members []discord.Member) string {
	var ids = []string{}

	for _, v := range members {
		ids = append(ids, fmt.Sprintf("<@!%s>", strconv.Itoa(int(v.User.ID))))
	}

	return strings.Join(ids, " ")

}

func chunkMembersArray(members []discord.Member) [][]discord.Member {
	var memsList = make([][]discord.Member, int(len(members)/50)+1)

	var i = 0
	for _, v := range members {
		if len(memsList) > 0 {
			if len(memsList[i]) == 50 {
				i++
			}
		}

		memsList[i] = append(memsList[i], v)
	}

	return memsList
}

// ListUnverified lists unverified users.
func (b *Bot) ListUnverified(c *gateway.MessageCreateEvent) (interface{}, error) {
	b.Ctx.Typing(c.ChannelID)

	GuildID := discord.GuildID(stuff.GuildID())

	members, err := b.Ctx.Members(GuildID)
	if err != nil {
		return e.FailedCommand("failed to get all members", err)
	}

	memsList := chunkMembersArray(b.getUnverifiedMembers(GuildID, members))

	if len(memsList) == 0 {
		return "No unverified users!", nil
	}

	for _, v := range memsList {
		mentions := strings.TrimSpace(JoinMemberMentions(v))

		if mentions == "" {
			break
		}

		b.Ctx.SendMessage(c.ChannelID, mentions, nil)
	}

	return nil, nil
}

// KickUnverified is a special command to kick members that hasn't verified yet.
func (b *Bot) KickUnverified(c *gateway.MessageCreateEvent) (interface{}, error) {
	b.Ctx.Typing(c.ChannelID)

	GuildID := discord.GuildID(stuff.GuildID())

	members, err := b.Ctx.Members(GuildID)
	if err != nil {
		return e.FailedCommand("failed to get all members", err)
	}

	b.handleKick(members, GuildID)

	return "Successfully removed all unverified members!", nil
}
