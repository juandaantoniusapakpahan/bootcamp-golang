package middleware

import (
	"fmt"
	"log"
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

func LogRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		fmt.Println(time.Now().Format(time.RFC3339), r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
