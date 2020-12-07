module server

go 1.15

require (
	github.com/impact-eintr/WEB-Project/Store_v2.0/HttpCache/cache v0.0.0
	github.com/impact-eintr/WEB-Project/Store_v2.0/HttpCache/http v0.0.0

)

replace (
	github.com/impact-eintr/WEB-Project/Store_v2.0/HttpCache/cache v0.0.0 => ../cache
	github.com/impact-eintr/WEB-Project/Store_v2.0/HttpCache/http v0.0.0 => ../http

)
