package main

import (
	"context"
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

	registerServices(mux, ctx)

	if err := http.ListenAndServe(
		":3000",
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Printf("failed to start server: %v", err)
		return
	}
}

func registerServices(mux *http.ServeMux, ctx context.Context) {

	application := service.NewApplication(ctx)

	login := ports.NewAuthGrpcServer(application)
	path, handler := pbv1connect.NewAuthServiceHandler(login)
	mux.Handle(path, handler)
}
