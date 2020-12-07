module server

go 1.15

require (
	github.com/HttpCache/cache v0.0.0
	github.com/HttpCache/http v0.0.0

)

replace (
	github.com/HttpCache/cache v0.0.0 => ../cache
	github.com/HttpCache/http v0.0.0 => ../http

)
