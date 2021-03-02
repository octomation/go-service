module service/tools

go 1.15

require (
	github.com/go-swagger/go-swagger v0.26.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.5.0
	github.com/golangci/golangci-lint v1.38.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.3.0
	github.com/kyoh86/looppointer v0.1.7
	github.com/twitchtv/twirp v7.1.1+incompatible
	golang.org/x/exp v0.0.0-20210220032938-85be41e4509f
	golang.org/x/tools v0.1.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.25.1-0.20201208041424-160c7477e0e8
)

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.1.0
