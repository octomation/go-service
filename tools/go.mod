module service/tools

go 1.15

require (
	github.com/go-swagger/go-swagger v0.25.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.3
	github.com/golangci/golangci-lint v1.34.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/kyoh86/looppointer v0.1.7
	github.com/twitchtv/twirp v7.1.0+incompatible
	golang.org/x/tools v0.5.0
)

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.5
