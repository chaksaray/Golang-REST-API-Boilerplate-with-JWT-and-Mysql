package middlewares

import (
	"net/http"
	"skeleton_project/app/auth"
	"skeleton_project/app/controllers"
)

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			controllers.RespondError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		next(w, r)
	}
}