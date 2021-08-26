package commands

import (
	"fmt"
	"strings"

	"github.com/World-of-Cryptopups/cordy/lib"
	"github.com/World-of-Cryptopups/cordy/lib/db"
	e "github.com/World-of-Cryptopups/cordy/lib/errors"
	rc "github.com/World-of-Cryptopups/cordy/lib/redis"
	"github.com/World-of-Cryptopups/cordy/stuff"
	"github.com/go-redis/redis/v8"

	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/enescakir/emoji"
)

func (b *Bot) Register(c *gateway.MessageCreateEvent, args bot.RawArguments) (string, error) {
	b.Ctx.Typing(c.ChannelID)

	// get discordid
	_discordId := c.Author.ID.String()

	// initialize deta db
	client, err := db.Client()
	if err != nil {
		return e.FailedCommand("error initializing deta db", err)
	}

	// check if user is already registered
	_registered, err := client.UserExists(_discordId)
	if err != nil {
		return e.FailedCommand("check if user is registered", err)
	}
	if _registered {
		return e.FailedMessage("You have already registered! If you want to change your acc, please contact an admin or mod.", err)
	}

	// get token
	if args == "" {
		return "", fmt.Errorf("%v No TOKEN provided", emoji.CrossMark)
	}
	token := strings.TrimSpace(string(args))

	// get initial datas from redis
	r := rc.Client()

	// check first if token / key exists from redis
	_e := r.Exists(rc.Ctx, "_token_"+token).Val()
	if _e == 0 {
		// key does not exist
		//lint:ignore ST1005 // I know what I am doing!
		return "", fmt.Errorf("I don't know that **TOKEN**, if you are not sure on what to do, please contact an admin or mod.")
	}

	val, err := r.HGetAll(rc.Ctx, "_token_"+token).Result()
	if err == redis.Nil {
		return e.FailedCommand("get all redis keys", err)
	}

	// check if token exists in fauna
	check, err := client.TokenExists(token)
	if err != nil {
		return e.FailedCommand("check if token exists already", err)
	}
	if check {
		return e.FailedMessage("This **TOKEN** has already been registered! If you did not register this, please contact an admin or mod.", err)
	}

	_wallet := val["wallet"]
	_type := val["type"]

	// confirm season pass info
	cfPass, err := stuff.ConfirmSeasonOnePass(_wallet)
	if err != nil {
		return e.FailedCommand("confirm seasonpass", err)
	}

	// fetch initial dps, call the function
	if d, err := stuff.FetchDPS(lib.UserDPSUser{
		ID:       c.Author.ID.String(),
		Username: c.Author.Tag(),
		Avatar:   c.Author.AvatarURL(),
	}, _wallet); err != nil {
		return e.FailedCommand("error in calling the api to get initial dps", err)
	} else {
		totalDPS := d.DPS.Pupcards + d.DPS.Pupskins + d.DPS.Pupitems.Real

		stuff.HandleUserRole(b.Ctx, discord.GuildID(stuff.GuildID()), int(c.Author.ID), totalDPS)
	}

	// fetch season pass details
	passDetails, err := stuff.GetSeasonOnePass(_wallet)
	if err != nil {
		return e.FailedCommand("get season one pass info", err)
	}

	// get current pass
	currentPass, err := stuff.GetCurrentPass(_wallet)
	if err != nil {
		fmt.Println(err)
	}

	// create user
	_user_ := &lib.User{
		Key: _discordId,
		User: lib.UserDiscord{
			ID:       c.Author.ID.String(),
			Username: c.Author.Username,
			Tag:      c.Author.Tag(),
			Avatar:   c.Author.AvatarURL(),
		},
		Wallet:      _wallet,
		Type:        _type,
		Token:       token,
		CurrentPass: currentPass.Pass,
		SeasonPasses: []lib.UserSeasonPass{{
			Season: cfPass.Season,
			Title:  cfPass.Pass,
			DPS:    passDetails.DPS,
		}},
	}

	// store data
	if _, err = client.DB.Put(*_user_); err != nil {
		return e.FailedCommand("create a new user", err)
	}

	return fmt.Sprintf("%v Successfully authenticated <@!%s>! You can now check your user info with `>me` command.", emoji.CheckBoxWithCheck, c.Author.ID), nil
}
