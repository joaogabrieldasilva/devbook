package middlewares

import (
	"api/src/authentication"
	"api/src/response"
	"log"
	"net/http"
)


func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if error := authentication.ValidateToken(r); error != nil {
			response.Error(w, http.StatusUnauthorized, error)
			return
		}
		next(w, r)
	}
}