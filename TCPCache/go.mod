module TCPCache

go 1.15

require (
	github.com/server/cache v0.0.0
	github.com/server/http v0.0.0
	github.com/server/tcp v0.0.0

)

replace (
	github.com/server/cache v0.0.0 => ./cache
	github.com/server/http v0.0.0 => ./http
	github.com/server/tcp v0.0.0 => ./tcp

)
