package ports

import (
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/gofrs/uuid/v5"

	pbv1 "github.com/migerick/cookie-authentication-go/protobuf/pb/v1"

	"github.com/migerick/cookie-authentication-go/internal/common/auth"
	"github.com/migerick/cookie-authentication-go/internal/users/app"
	"github.com/migerick/cookie-authentication-go/internal/users/app/command"
	"github.com/migerick/cookie-authentication-go/internal/users/app/query"
	"github.com/migerick/cookie-authentication-go/protobuf/pb/v1/pbv1connect"
)

type AuthGrpcServer struct {
	app app.Application
}

func (a AuthGrpcServer) GetUsers(ctx context.Context, _ *connect.Request[pbv1.GetUsersRequest]) (*connect.Response[pbv1.GetUsersResponse], error) {
	response, err := a.app.Queries.GetUsers.Handle(ctx, query.User{Query: "test"})
	if err != nil {
		return nil, err
	}

	return connect.NewResponse(&pbv1.GetUsersResponse{
		Users: response,
	}), nil
}

func (a AuthGrpcServer) Logout(_ context.Context, _ *connect.Request[pbv1.LogoutRequest]) (*connect.Response[pbv1.LogoutResponse], error) {
	response := connect.NewResponse(&pbv1.LogoutResponse{
		Success: true,
		Message: "Logout successful",
	})

	cookie := &auth.Cookie[pbv1.LogoutResponse]{}
	cookie.Delete(response)

	return response, nil
}

func (a AuthGrpcServer) Login(ctx context.Context, c *connect.Request[pbv1.LoginRequest]) (*connect.Response[pbv1.LoginResponse], error) {
	if err := a.app.Commands.Login.Handle(ctx, command.Login{Email: c.Msg.Email, Password: c.Msg.Password}); err != nil {
		return nil, err
	}

	response := connect.NewResponse(&pbv1.LoginResponse{
		Success: true,
		Message: "Login successful",
	})

	cookie := &auth.Cookie[pbv1.LoginResponse]{}
	cookie.Set(response, uuid.NewV5(uuid.NamespaceOID, c.Msg.Email).String())

	return response, nil
}

var _ pbv1connect.AuthServiceHandler = &AuthGrpcServer{}

func NewAuthGrpcServer(app app.Application) *AuthGrpcServer {
	return &AuthGrpcServer{app: app}
}
