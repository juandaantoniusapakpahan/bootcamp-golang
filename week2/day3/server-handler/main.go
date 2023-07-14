package main

import (
	"fmt"
	"net/http"
)

type Hander1 struct {
}

func (h1 *Hander1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, map[string]interface{}{"status": "success", "data": "Handler1"})
}

type Handler2 struct {
}

func (h2 *Handler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, map[string]interface{}{"status": "success", "data": "Handler2"})
}

func main() {
	handler1 := Hander1{}
	handler2 := Handler2{}

	http.Handle("/handle1", &handler1)
	http.Handle("/handle2", &handler2)

	fmt.Println("Server is strated at localhost:8080")
	http.ListenAndServe(":8080", nil)

}
