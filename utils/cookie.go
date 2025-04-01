package utils

import (
	"net/http"
	"time"
)

func GenerateCookie(token string) http.Cookie {
	cookie := http.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 6),
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	}

	return cookie
}
