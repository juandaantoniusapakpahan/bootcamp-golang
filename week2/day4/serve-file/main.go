package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./assets"))))
	mux.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./assets")
	})
	mux.HandleFunc("/kosong/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
