package auth

import (
	"github.com/bufbuild/connect-go"
	"net/http"
	"time"
)

type Cookie[T any] struct{}

func NewCookie[T any]() Cookie[T] {
	return Cookie[T]{}
}

func (c Cookie[T]) Set(request *connect.Response[T], token string) {
	cookie := SetCookie(token)
	request.Header().Set("Set-Cookie", cookie.String())
}

func (c Cookie[T]) Delete(request *connect.Response[T]) {
	cookie := SetCookie("")
	request.Header().Set("Set-Cookie", cookie.String())
}

func SetCookie(token string) http.Cookie {
	cookie := http.Cookie{
		Name:     "auth_cookie",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
	}

	//TODO: Uncomment this when we have https
	//cookie.Secure = true

	if token == "" {
		// Set cookie to expire immediately
		cookie.Expires = time.Now().Add(0 * time.Second)
		cookie.MaxAge = -1
		return cookie
	}

	// Set cookie to expire in 1 hour
	cookie.Expires = time.Now().Add(1 * time.Hour)
	cookie.MaxAge = 3600
	return cookie
}
