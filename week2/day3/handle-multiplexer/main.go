package main

import (
	"fmt"
	"net/http"
)

type Handler1 struct {
}

func (h1 *Handler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandlerMux")
}

func main() {
	mux := http.NewServeMux()
	handler1 := Handler1{}

	mux.Handle("/handler", &handler1)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server started at localhost:", server.Addr)
	server.ListenAndServe()
}
