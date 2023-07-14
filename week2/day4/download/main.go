package main

import (
	"io"
	"net/http"
	"os"
)

func Download(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./assets/test.txt")
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+file.Name())
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	io.Copy(w, file)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/download", Download)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
