package server

import (
	"context"
	"fmt"

	"go.octolab.org/template/service"
	v1 "go.octolab.org/template/service/api/rpc/v1"
)

type GRPC struct {
	v1.UnimplementedGreeterServiceServer
}

func (srv *GRPC) Hello(
	_ context.Context,
	req *v1.HelloRequest,
) (*v1.HelloResponse, error) {
	return &v1.HelloResponse{
		Text: fmt.Sprintf("Hello, %s!", req.GetSubject()),
	}, nil
}

func (srv *GRPC) Sign(
	context.Context,
	*v1.SignRequest,
) (*v1.SignResponse, error) {
	return &v1.SignResponse{License: service.License}, nil
}
