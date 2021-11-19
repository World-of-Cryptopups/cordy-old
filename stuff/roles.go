package stuff

import (
	"fmt"
	"os"
	"strings"

	"github.com/World-of-Cryptopups/cordy/utils"
	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
)

type DPSStats struct {
	Title  string         // Title is the name of the Role
	RoleID discord.RoleID // RoleID of the Role
	Color  string         // Color of the Role
}

var initRoles = strings.Split(os.Getenv("ROLES"), ",")

const PupsWarrior = 3000
const PupsKnight = 5000
const PupsOverlord = 8000
const PupsApocalypse = 10000
const PupsAboveAll = 25000
const PupsDoggosOfInfinity = 70000
const PupsDoggosOfEternity = 145000

// Roles is the roles and
var Roles = map[int]DPSStats{
	PupsWarrior: {
		Title:  "Warrior Pups",
		RoleID: discord.RoleID(utils.ConvertInt(initRoles[0])), // change this IDs in production
		Color:  "green",
	},
	PupsKnight: {
		Title:  "Knight Pups",
		RoleID: discord.RoleID(utils.ConvertInt(initRoles[1])), // change this IDs in production
		Color:  "blue",
	},
	PupsOverlord: {
		Title:  "Overlord Pups",
		RoleID: discord.RoleID(utils.ConvertInt(initRoles[2])), // change this IDs in production
		Color:  "purple",
	},
	PupsApocalypse: {
		Title:  "Pups of the Apocalypse",
		RoleID: discord.RoleID(utils.ConvertInt(initRoles[3])), // change this IDs in production
		Color:  "red",
	},
	PupsAboveAll: {
		Title:  "Pups Above All",
		RoleID: discord.RoleID(utils.ConvertInt(initRoles[4])), // change this IDs in production
		Color:  "orange",
	},
	PupsDoggosOfInfinity: {
		Title:  "Doggos of Infinity",
		RoleID: discord.RoleID(utils.ConvertInt(initRoles[5])), // change this IDs in production
		Color:  "gold",
	},
	PupsDoggosOfEternity: {
		Title:  "Doggos of Eternity",
		RoleID: discord.RoleID(utils.ConvertInt(initRoles[6])), // change this IDs in production
		Color:  "white",
	},
}

// AllRoles is the list of all available roles for ranking.
var AllRoles = []string{
	"Warrior Pups",
	"Knight Pups",
	"Overlord Pups",
	"Pups of the Apocalypse",
	"Pups Above All",
	"Doggos of Infinity",
	"Doggos of Eternity",
}

var Colors = map[string]string{
	"purple": "#a652bb",
	"blue":   "#3b6fff",
	"cyan":   "#00c09a",
	"green":  "#00d166",
	"gold":   "#fff000",
	"red":    "#e61616",
	"orange": "#ffa500",
	"white":  "#ffffff",
	"grey":   "#95a5a6",
}

// GetHighestRole gets the member's current role.
func GetHighestRole(member *discord.Member) DPSStats {
	var d DPSStats

	for _, v := range Roles {
		for _, x := range member.RoleIDs {
			if v.RoleID == x {
				d = v
			}
		}
	}

	return d
}

// GetDPSRoleInfo gets the role info for a specific DPS.
func GetDPSRoleInfo(dps int) DPSStats {
	var d DPSStats

	if dps >= PupsWarrior && dps < PupsKnight {
		d = Roles[PupsWarrior]
	} else if dps >= PupsKnight && dps < PupsOverlord {
		d = Roles[PupsKnight]
	} else if dps >= PupsOverlord && dps < PupsApocalypse {
		d = Roles[PupsOverlord]
	} else if dps >= PupsApocalypse && dps < PupsAboveAll {
		d = Roles[PupsApocalypse]
	} else if dps >= PupsAboveAll && dps < PupsDoggosOfInfinity {
		d = Roles[PupsAboveAll]
	} else if dps >= PupsDoggosOfInfinity && dps < PupsDoggosOfEternity {
		d = Roles[PupsDoggosOfInfinity]
	} else if dps >= PupsDoggosOfEternity {
		d = Roles[PupsDoggosOfEternity]
	}

	return d
}

// HandleUserRole handles the user's role, could remove the old one and change it.
func HandleUserRole(ctx *bot.Context, guildID discord.GuildID, discordID int, dps int) error {
	fmt.Println(guildID)
	member, err := ctx.Client.Member(guildID, discord.UserID(discordID))
	if err != nil {
		fmt.Println("Error in getting the member in the guild")
		return err
	}

	d := GetDPSRoleInfo(dps)
	if d.Title != "" {
		// promote user
		PromoteUser(ctx, guildID, member.User.ID, dps)
	}

	return nil
}

// PromoteUser handles adding of role from base to lowest.
func PromoteUser(ctx *bot.Context, guildID discord.GuildID, userid discord.UserID, dps int) {
	for i, v := range Roles {
		if i <= dps {
			// if role is lower or equal to the promote role, add it
			ctx.Client.AddRole(guildID, userid, v.RoleID)
		} else {
			// otherwise, remove the role from user
			ctx.Client.RemoveRole(guildID, userid, v.RoleID)
		}
	}
}
