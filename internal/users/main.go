package main

import (
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/migerick/cookie-authentication-go/internal/users/interceptor"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/migerick/cookie-authentication-go/internal/users/ports"
	"github.com/migerick/cookie-authentication-go/internal/users/service"
	"github.com/migerick/cookie-authentication-go/protobuf/pb/v1/pbv1connect"
)

func main() {
	ctx := context.Background()

	mux := http.NewServeMux()

	interceptors := connect.WithInterceptors(
		interceptor.NewAuthInterceptor(),
	)

	registerServices(mux, ctx, interceptors)

	if err := http.ListenAndServe(
		":3000",
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Printf("failed to start server: %v", err)
		return
	}
}

func registerServices(mux *http.ServeMux, ctx context.Context, interceptors connect.Option) {

	application := service.NewApplication(ctx)

	login := ports.NewAuthGrpcServer(application)
	path, handler := pbv1connect.NewAuthServiceHandler(login, interceptors)
	mux.Handle(path, handler)
}
