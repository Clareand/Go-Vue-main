package middlewares

import (
	"errors"
	"net/http"

	"github.com/Clareand/rest-api/api/auth"
	"github.com/Clareand/rest-api/api/helpers"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r)
	}
}

func setMiddlewareAuthenticaation(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		err := auth.TokenValid(r)
		if err != nil {
			helpers.Error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
