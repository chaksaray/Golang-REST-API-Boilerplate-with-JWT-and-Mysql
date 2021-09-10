package middlewares

import (
	"net/http"

	"github.com/chaksaray/Golang-REST-API-Boilerplate-with-JWT-and-Mysql/app/auth"
	"github.com/chaksaray/Golang-REST-API-Boilerplate-with-JWT-and-Mysql/app/controllers"
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