package common

type Option int

const (
	_ Option = iota
	LOGIN
	SIGNIN
	EXIT
)

type User struct {
	Uid string
	Pwd string
}
