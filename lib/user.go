package lib

//
type User struct {
	Key         string      `json:"key"` // this is needed by deta
	User        UserDiscord `json:"user"`
	Wallet      string      `json:"wallet"`
	Type        string      `json:"type"`
	Token       string      `json:"token"`
	CurrentPass string      `json:"currentPass"`
	Rank        int         `json:"rank"`
}

type UserDiscord struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Tag      string `json:"tag"`
}
