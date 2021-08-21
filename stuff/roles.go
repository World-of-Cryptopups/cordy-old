package stuff

import (
	"fmt"

	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
)

type DPSStats struct {
	Title  string         // Title is the name of the Role
	RoleID discord.RoleID // RoleID of the Role
	Color  string         // Color of the Role
}

// Roles is the roles and
var Roles = map[int]DPSStats{
	3000: {
		Title:  "Warrior Pups",
		RoleID: 878236495926919208, // change this IDs in production
		Color:  "green",
	},
	5000: {
		Title:  "Knight Pups",
		RoleID: 878236594807668767, // change this IDs in production
		Color:  "blue",
	},
	8000: {
		Title:  "Overlord Pups",
		RoleID: 878236658699497483, // change this IDs in production
		Color:  "purple",
	},
	10000: {
		Title:  "Pups of the Apocalypse",
		RoleID: 878236785979822151, // change this IDs in production
		Color:  "red",
	},
	20000: {
		Title:  "Pups Above All",
		RoleID: 878236838698045520, // change this IDs in production
		Color:  "orange",
	},
	100000: {
		Title:  "Doggos of Infinity",
		RoleID: 878236874643226675, // change this IDs in production
		Color:  "gold",
	},
	200000: {
		Title:  "Doggos of Eternity",
		RoleID: 878236916900827167, // change this IDs in production
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
	"violet": "#9b59b6",
	"blue":   "#3498db",
	"cyan":   "#2ecc71",
	"green":  "#2ecc71",
	"gold":   "#e67e22",
	"red":    "#e91e63",
	"orange": "#e74c3c",
	"white":  "#ffffff",
	"grey":   "#95a5a6",
}

// HasCurrentRole gets the member's current role.
func HasCurrentRole(member *discord.Member) (DPSStats, bool) {
	for _, v := range Roles {
		for _, x := range member.RoleIDs {
			if v.RoleID == x {
				return v, true
			}
		}
	}

	return DPSStats{}, false
}

// GetDPSRoleInfo gets the role info for a specific DPS.
func GetDPSRoleInfo(dps int) DPSStats {
	var d DPSStats

	if dps >= 3000 && dps < 5000 {
		d = Roles[3000]
	} else if dps >= 5000 && dps < 8000 {
		d = Roles[5000]
	} else if dps >= 8000 && dps < 10000 {
		d = Roles[8000]
	} else if dps >= 10000 && dps < 20000 {
		d = Roles[10000]
	} else if dps >= 20000 && dps < 100000 {
		d = Roles[20000]
	} else if dps >= 100000 && dps < 200000 {
		d = Roles[100000]
	} else if dps >= 200000 {
		d = Roles[200000]
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
	fmt.Println(member)

	d := GetDPSRoleInfo(dps)
	fmt.Println(d)
	if d.Title != "" {
		currentRole, _ := HasCurrentRole(member)
		if currentRole.Title != d.Title {
			ctx.Client.RemoveRole(guildID, member.User.ID, currentRole.RoleID)
			ctx.Client.AddRole(guildID, member.User.ID, d.RoleID)
		}
	}

	return nil
}
