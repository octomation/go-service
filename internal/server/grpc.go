package server

import (
	"context"
	"fmt"

	"go.octolab.org/template/service"
	servicev1 "go.octolab.org/template/service/internal/api/service/v1"
)

type GRPC struct {
	servicev1.UnimplementedGreeterServiceServer
}

func (srv *GRPC) Hello(
	_ context.Context,
	req *servicev1.HelloRequest,
) (*servicev1.HelloResponse, error) {
	return &servicev1.HelloResponse{
		Text: fmt.Sprintf("Hello, %s!", req.GetSubject()),
	}, nil
}

func (srv *GRPC) Sign(
	context.Context,
	*servicev1.SignRequest,
) (*servicev1.SignResponse, error) {
	return &servicev1.SignResponse{License: service.License}, nil
}
