package commands

import (
	"fmt"
	"strings"

	e "github.com/World-of-Cryptopups/roleroll-new/lib/errors"
	fc "github.com/World-of-Cryptopups/roleroll-new/lib/fauna"
	rc "github.com/World-of-Cryptopups/roleroll-new/lib/redis"
	"github.com/World-of-Cryptopups/roleroll-new/stuff"
	"github.com/go-redis/redis/v8"

	"github.com/diamondburned/arikawa/v2/bot"
	"github.com/diamondburned/arikawa/v2/gateway"
	"github.com/enescakir/emoji"
	f "github.com/fauna/faunadb-go/v4/faunadb"
)

func (b *Bot) Register(c *gateway.MessageCreateEvent, args bot.RawArguments) (string, error) {
	if args == "" {
		return "", fmt.Errorf("%v No TOKEN provided", emoji.CrossMark)
	}

	// get discordid
	_discordId := c.Author.ID.String()

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

	// fauna client
	fauna := fc.Client()

	// check if user is already registered
	_registered, err := fc.CheckUser(_discordId)
	if err != nil {
		return e.FailedCommand("check if user is registered", err)
	}
	if _registered {
		return e.FailedMessage("You have already registered! If you want to change your acc, please contact an admin or mod.", err)
	}

	// check if token exists in fauna
	check, err := fauna.Query(f.Exists(f.MatchTerm(f.Index("userByToken"), token)))
	if err != nil {
		return e.FailedCommand("check if token exists already", err)
	}
	var ccc bool
	check.Get(&ccc)
	if ccc {
		return e.FailedMessage("This **TOKEN** has already been registered! If you did not register this, please contact an admin or mod.", err)
	}

	_wallet := val["wallet"]
	_type := val["type"]

	// confirm season pass info
	cfPass, err := stuff.ConfirmSeasonOnePass(_wallet)
	if err != nil {
		return e.FailedCommand("confirm seasonpass", err)
	}

	// create user
	_user_ := User{
		DiscordID:       c.Author.ID.String(),
		DiscordUsername: c.Author.Username,
		AvatarURL:       c.Author.AvatarURL(),
		Wallets:         []string{_wallet},
		DefaultWallet:   _wallet,
		Type:            _type,
		Token:           token,
		SeasonPasses:    []UserSeasonPasses{{Season: cfPass.Season, Title: cfPass.Pass}},
	}
	user, err := fauna.Query(f.Create(f.Collection("users"), f.Obj{"data": _user_}))
	if err != nil {
		return e.FailedCommand("create a new user", err)
	}

	// fetch season pass details
	passDetails, err := stuff.GetSeasonOnePass(_wallet)
	if err != nil {
		return e.FailedCommand("get season one pass info", err)
	}

	var userRef f.RefV
	user.At(f.ObjKey("ref")).Get(&userRef)

	// create a new season pass document
	_userPass_ := UserSeasonPass{
		User: userRef,
		DPS:  passDetails.DPS,
	}
	if _, err = fauna.Query(f.Create(f.Collection("seasonpass"), f.Obj{"data": _userPass_})); err != nil {
		return e.FailedCommand("create season pass document", err)
	}

	return fmt.Sprintf("%v Successfully authenticated <@!%s>!", emoji.CheckBoxWithCheck, c.Author.ID), nil
}
