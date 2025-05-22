package middlewares

import (
	"context"
	"net/http"

	"github.com/ivanglzr/PaintBackend/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path == "/auth/log-in" || path == "auth/register" {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie("access_token")

		if err != nil {
			if err == http.ErrNoCookie {
				utils.JSONResponse(w, 401, "Petition unathorized")
				return
			}

			utils.JSONResponse(w, 500, "Error while reading the cookie")
			return
		}

		id, err := utils.DecodeToken(cookie.Value)

		if err != nil {
			utils.JSONResponse(w, 401, "Petition unathorized")
			return
		}

		ctx := context.WithValue(r.Context(), "id", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
