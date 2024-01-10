package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/migerick/cookie-authentication-go/config"
	"github.com/migerick/cookie-authentication-go/handler"
)

func main() {
	cfg := config.Parsed

	//if err := database.NewConnectionPostgres(cfg.PostgresURL); err != nil {
	//	throw(err)
	//}

	mux := http.NewServeMux()
	mux.Handle("/", handler.Muxer)

	opts := cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedHeaders: []string{
			// standard
			"Accept-Encoding", "Content-Encoding", "Content-Type",
			// Connect
			"Connect-Protocol-Version", "Connect-Timeout-Ms",
			// future Connect
			"Connect-Accept-Encoding", "Connect-Content-Encoding",
			// gRPC-web
			"Grpc-Timeout", "X-Grpc-Web", "X-User-Agent",
		},
		ExposedHeaders: []string{
			// future Connect
			"Content-Encoding", "Connect-Content-Encoding",
			// gRPC-web
			"Grpc-Status", "Grpc-Message",
		},
		AllowCredentials: true,
	}
	if cfg.Debug {
		opts.AllowOriginFunc = func(string) bool { return true }
	} else {
		opts.AllowedOrigins = strings.Split(cfg.Origins, ",")
	}

	lis := must(net.Listen("tcp", cfg.Addr))
	server := &http.Server{Handler: h2c.NewHandler(cors.New(opts).Handler(mux), &http2.Server{})}
	defer server.Close()

	fmt.Printf("Middle Service Starts | %s\n", lis.Addr())
	if err := server.Serve(lis); !errors.Is(err, http.ErrServerClosed) {
		throw(err)
	}
}

func throw(err error) {
	if err != nil {
		panic(err)
	}
}

func must[T any](val T, err error) T {
	throw(err)
	return val
}
