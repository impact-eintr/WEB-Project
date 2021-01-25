package common

type Option int

const (
	_ Option = iota
	LOGIN
	LOGUP
	EXIT
)

type User struct {
	Uid string
	Pwd string
}
