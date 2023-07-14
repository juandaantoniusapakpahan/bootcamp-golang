package main

import (
	"fmt"
	"net/http"
	"time"
)

func SetCookied(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newCookie := &http.Cookie{
			Name:    "TestCookies",
			Value:   "Icanhandlethisforyou",
			Path:    "/",
			MaxAge:  3 * 60,
			Expires: time.Now().Add(3 * time.Minute),
		}
		http.SetCookie(w, newCookie)
		next.ServeHTTP(w, r)
	})

}

func Handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Handler1: Check your cookies")
}
func Handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Handler2: Check your cookies")
}

func main() {
	mux := http.NewServeMux()
	finalHandler := http.HandlerFunc(Handler1)

	mux.Handle("/handler1", SetCookied(finalHandler))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
