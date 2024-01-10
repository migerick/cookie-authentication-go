package handler

import (
	"net/http"

	"connectrpc.com/connect"
)

var Muxer = &muxer{http.NewServeMux()}

type muxer struct {
	serverMux *http.ServeMux
}

func (m *muxer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.serverMux.ServeHTTP(w, r)
}

func (m *muxer) addHandlerToMux(pattern string, handler http.Handler) {
	m.serverMux.Handle(pattern, handler)
}

func addGrpc[T any](connectFunc func(T, ...connect.HandlerOption) (string, http.Handler), svc any, opts ...connect.HandlerOption) {
	defaultOpts := []connect.HandlerOption{
		connect.WithInterceptors(),
	}
	pattern, handler := connectFunc(svc.(T), append(defaultOpts, opts...)...)
	Muxer.addHandlerToMux(pattern, handler)
}
