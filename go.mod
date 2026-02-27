module github.com/lucas59356/notify

go 1.24.3

require (
	github.com/lucas59356/go-logger v0.0.0-00010101000000-000000000000
	github.com/mattn/go-gntp v0.0.0-20200109101910-88aa4ee0ab11
	github.com/urfave/cli v1.22.17
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.7 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
)

replace github.com/lucas59356/go-logger => ./local/go-logger
