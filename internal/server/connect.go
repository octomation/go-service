package server

import (
	"context"
	"fmt"

	"connectrpc.com/connect"

	"go.octolab.org/template/service"
	v1 "go.octolab.org/template/service/api/rpc/v1"
	"go.octolab.org/template/service/api/rpc/v1/v1connect"
)

type Connect struct {
	v1connect.UnimplementedGreeterServiceHandler
}

func (srv *Connect) Hello(
	_ context.Context,
	req *connect.Request[v1.HelloRequest],
) (*connect.Response[v1.HelloResponse], error) {
	return connect.NewResponse(&v1.HelloResponse{
		Text: fmt.Sprintf("Hello, %s!", req.Msg.GetSubject()),
	}), nil
}

func (srv *Connect) Sign(
	context.Context,
	*connect.Request[v1.SignRequest],
) (*connect.Response[v1.SignResponse], error) {
	return connect.NewResponse(&v1.SignResponse{
		License: service.License,
	}), nil
}
