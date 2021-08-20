package stuff

import (
	"os"
	"strconv"
)

// GuildID is the servers' id, (CHANGE THIS IN PRODUCTION)
func GuildID() int {
	var g, _ = strconv.Atoi(os.Getenv("DEFAULT_GUILDID"))
	return g
}
