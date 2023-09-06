package auth

import (
	"github.com/bufbuild/connect-go"
	"net/http"
)

type Cookie[T any] struct{}

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
		cookie.MaxAge = -1
		return cookie
	}

	// Set cookie to expire in 1 hour
	cookie.MaxAge = 3600
	return cookie
}
