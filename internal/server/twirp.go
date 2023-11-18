package server

import v1 "go.octolab.org/template/service/api/rpc/v1"

func Twirp() v1.TwirpServer {
	return v1.NewGreeterServiceServer(new(GRPC))
}
