package main

import (
	"fmt"
	"net/http"
)

type FirstHandler struct {
}

func (f *FirstHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

var Handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Success")
}

func main() {
	// handler := &FirstHandler{}
	server := http.Server{
		Addr:    ":8080",
		Handler: Handler,
	}
	fmt.Printf("Server running on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
