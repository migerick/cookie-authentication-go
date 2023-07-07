package interceptor

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	"github.com/migerick/cookie-authentication-go/protobuf/pb/v1/pbv1connect"
)

type AuthInterceptor struct {
}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

func (interceptor *AuthInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(
		ctx context.Context,
		req connect.AnyRequest,
	) (connect.AnyResponse, error) {
		if shouldSkipAuth(req.Spec().Procedure) {
			return next(ctx, req)
		}

		if req.Header().Get("Cookie") == "" {
			return nil, fmt.Errorf("cookie is empty, please login")
		}

		return next(ctx, req)
	}
}

func shouldSkipAuth(method string) bool {
	switch method {
	case pathServiceAuth("Login"), pathServiceAuth("Logout"):
		return true
	default:
		return false
	}
}

func (interceptor *AuthInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(
		ctx context.Context,
		conn connect.StreamingHandlerConn,
	) error {
		if conn.RequestHeader().Get("Cookie") == "" {
			return fmt.Errorf("cookie is empty, please login")
		}
		return next(ctx, conn)
	}
}

func (*AuthInterceptor) WrapStreamingClient(connect.StreamingClientFunc) connect.StreamingClientFunc {
	panic("not implemented")
}

func pathServiceAuth(path string) string {
	return fmt.Sprintf("/%s/%s", pbv1connect.AuthServiceName, path)
}
