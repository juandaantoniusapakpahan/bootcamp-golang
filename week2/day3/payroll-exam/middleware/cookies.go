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
		// r.RemoteAddr:  to get the remote client's IP address and port
		fmt.Println(time.Now().Format(time.DateTime), r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

type CustomMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		current = next(current)
	}
	current.ServeHTTP(w, r)
}
