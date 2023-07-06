package ports

import (
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/migerick/cookie-authentication-go/internal/users/app/command"

	"github.com/migerick/cookie-authentication-go/internal/users/app"
	"github.com/migerick/cookie-authentication-go/protobuf/pb/v1/pbv1connect"

	pbv1 "github.com/migerick/cookie-authentication-go/protobuf/pb/v1"
)

type AuthGrpcServer struct {
	app app.Application
}

func (a AuthGrpcServer) Logout(ctx context.Context, c *connect.Request[pbv1.LogoutRequest]) (*connect.Response[pbv1.LogoutResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthGrpcServer) Login(ctx context.Context, c *connect.Request[pbv1.LoginRequest]) (*connect.Response[pbv1.LoginResponse], error) {
	if err := a.app.Commands.Login.Handle(ctx, command.Login{Email: c.Msg.Email, Password: c.Msg.Password}); err != nil {
		return nil, err
	}

	return connect.NewResponse(&pbv1.LoginResponse{
		Success: true,
		Message: "Login successful",
	}), nil
}

var _ pbv1connect.AuthServiceHandler = &AuthGrpcServer{}

func NewAuthGrpcServer(app app.Application) *AuthGrpcServer {
	return &AuthGrpcServer{app: app}
}
