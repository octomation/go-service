module service/tools

go 1.13

require (
	github.com/go-swagger/go-swagger v0.25.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.2
	github.com/golangci/golangci-lint v1.30.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/kyoh86/looppointer v0.1.6
	github.com/twitchtv/twirp v5.12.1+incompatible
	golang.org/x/tools v0.3.3
)

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.3
