package main

import (
	"fmt"
	"net/http"
	"time"
)

func HandlerFunc1(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "GG",
		Value:    "ddg",
		MaxAge:   int(time.Now().Unix()),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	fmt.Fprint(w, "GGWP")
}
func HandlerFunc2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandlerFunc1")
}

func main() {

	http.HandleFunc("/handlerfunc1", HandlerFunc1)
	http.HandleFunc("/handlerfunc2", HandlerFunc2)

	http.ListenAndServe(":8080", nil)
}
