package main

import (
	"fmt"
	"net/http"
)

// Handler

type HelloStruct struct {
}

func (h *HelloStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Helo Guys")
}

type WorldStruct struct {
}

func (wo *WorldStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "World is mine")
}

// HandlerFunc
func TestHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HandlerFunc")
}

func TestHandlerFunc2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Handler Func 2")
}

func main() {
	hello := HelloStruct{}
	world := WorldStruct{}

	// HandlerFunc
	var handlerFunc http.HandlerFunc = TestHandlerFunc

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	http.HandleFunc("/handlerfunc", handlerFunc)
	http.HandleFunc("/handler/2", TestHandlerFunc2)

	fmt.Println("Server running on localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

}
