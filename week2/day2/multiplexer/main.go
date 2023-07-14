package main

import (
	"fmt"
	"net/http"
)

// http.Handler
type HandleFuncI struct {
}

func (h *HandleFuncI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandlerI")
}

type HandleFuncII struct {
}

func (h *HandleFuncII) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandlerII")
}

// http.HandlerFunc
func HandlerFuncI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandlerFuncI")
}
func HandlerFuncII(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandlerFuncII")

}

func main() {
	mux := http.NewServeMux()

	handleFuncI := &HandleFuncI{}
	handleFuncII := &HandleFuncII{}

	mux.Handle("/handler1", handleFuncI)
	mux.Handle("/handler2", handleFuncII)
	mux.HandleFunc("/handlerfunc1", HandlerFuncI)
	mux.HandleFunc("/handlerfunc2", HandlerFuncII)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Printf("Server running on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
