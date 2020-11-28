module service/tools

go 1.15

require (
	github.com/go-swagger/go-swagger v0.25.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.4.4
	github.com/golangci/golangci-lint v1.35.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.1.0
	github.com/kyoh86/looppointer v0.1.7
	github.com/twitchtv/twirp v7.1.0+incompatible
	golang.org/x/exp v0.0.0-20201229011636-eab1b5eb1a03
	golang.org/x/tools v0.5.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.25.0
)

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.5
