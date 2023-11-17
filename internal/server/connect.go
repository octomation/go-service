package server

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"go.octolab.org/template/service"
	servicev1 "go.octolab.org/template/service/internal/api/service/v1"
	"go.octolab.org/template/service/internal/api/service/v1/servicev1connect"
)

type Connect struct {
	servicev1connect.UnimplementedGreeterServiceHandler
}

func (srv *Connect) Hello(
	_ context.Context,
	req *connect.Request[servicev1.HelloRequest],
) (*connect.Response[servicev1.HelloResponse], error) {
	return connect.NewResponse(&servicev1.HelloResponse{
		Text: fmt.Sprintf("Hello, %s!", req.Msg.GetSubject()),
	}), nil
}

func (srv *Connect) Sign(
	context.Context,
	*connect.Request[servicev1.SignRequest],
) (*connect.Response[servicev1.SignResponse], error) {
	return connect.NewResponse(&servicev1.SignResponse{
		License: service.License,
	}), nil
}
