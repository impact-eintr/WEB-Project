package common

type Option int

const (
	_ Option = iota
	LOGIN
	SIGNIN
	EXIT
)

type Status int

const (
	_ Status = iota
	ONLINE
	OFFLINE
)

type User struct {
	Uid     string `json:"uid"`
	Pwd     string `json:"pwd"`
	Uname   string `json:"uname"`
	Ustatus Status `json:"ustatus"`
}
