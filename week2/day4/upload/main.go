package main

import (
	"io"
	"net/http"
	"os"
)

func MultiUpload(w http.ResponseWriter, r *http.Request) {
	multipart, err := r.MultipartReader()
	if err != nil {
		panic(err)
	}

	for {
		part, err := multipart.NextPart()
		if err == io.EOF {
			break
		}
		defer part.Close()
		file, err := os.Create("assets/" + part.FileName())
		if err != nil {
			panic(err)
		}
		defer part.Close()
		_, err = io.Copy(file, part)
		if err != nil {
			panic(err)
		}

	}
	w.Write([]byte("GGWP"))
}

func Upload(w http.ResponseWriter, r *http.Request) {
	multipart, fileheader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	defer multipart.Close()

	file, err := os.Create("assets/" + fileheader.Filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, multipart)
	if err != nil {
		panic(err)
	}

	w.Write([]byte("GGWP"))
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/upload", Upload)
	mux.HandleFunc("/multiupload", MultiUpload)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
