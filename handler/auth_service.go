package handler

import (
	"context"

	"connectrpc.com/connect"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/migerick/cookie-authentication-go/gen/go/pb/v1"
	"github.com/migerick/cookie-authentication-go/gen/go/pb/v1/pbv1connect"
)

func init() {
	addGrpc(pbv1connect.NewAuthServiceHandler, &AuthHandler{})
}

type AuthHandler struct {
}

func (a AuthHandler) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[emptypb.Empty], error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthHandler) Logout(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[emptypb.Empty], error) {
	//TODO implement me
	panic("implement me")
}

var _ pbv1connect.AuthServiceHandler = (*AuthHandler)(nil)
