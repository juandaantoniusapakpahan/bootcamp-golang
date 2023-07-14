package middleware

import (
	"net/http"
	"time"
)

func SetMyCookies(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() != "/cookies" {
			newCookie := &http.Cookie{
				Name:    "PayRoll",
				Value:   "ajsdoifhasfh9w8eufasefa",
				MaxAge:  4 * 60,
				Path:    "/",
				Expires: time.Now().Add(4 * time.Minute),
			}

			http.SetCookie(w, newCookie)
		}

		next.ServeHTTP(w, r)
	})

}
