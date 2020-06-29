module github.com/benchhub/benchhub

go 1.14

require (
	github.com/dyweb/gommon v0.0.13
	github.com/gogo/protobuf v1.3.1
	golang.org/x/sys v0.0.0-20190422165155-953cdadca894 // indirect
	google.golang.org/grpc v1.27.1
)

replace github.com/dyweb/gommon => ../../dyweb/gommon
