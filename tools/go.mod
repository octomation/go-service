module service/tools

go 1.16

require (
	github.com/go-swagger/go-swagger v0.27.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/golangci/golangci-lint v1.41.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	github.com/twitchtv/twirp v8.1.0+incompatible
	golang.org/x/exp v0.0.0-20210618142145-ffcf9a09ea36
	golang.org/x/tools v0.1.3
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.26.0
)

replace (
	github.com/golangci/golangci-lint => github.com/kamilsk/golangci-lint v1.41.1
	golang.org/x/tools => github.com/kamilsk/go-tools v0.1.3
)
