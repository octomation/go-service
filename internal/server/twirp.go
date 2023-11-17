package server

import servicev1 "go.octolab.org/template/service/internal/api/service/v1"

func Twirp() servicev1.TwirpServer {
	return servicev1.NewGreeterServiceServer(new(GRPC))
}
