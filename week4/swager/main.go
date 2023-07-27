package main

import (
	"net/http"
	"swagger/helper"
)

type Category struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
type CategoryHandler struct {
}

func (c *Category) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		category := Category{
			FirstName: "Juanda Antonius ",
			LastName:  "Pakpahan",
		}
		response := helper.ResponseTemplete{
			Code:   http.StatusOK,
			Status: "success",
			Data:   map[string]interface{}{"category": category},
		}
		helper.WriteToResponse(w, response, response.Code)

	default:
		response := helper.ResponseTemplete{
			Code:   http.StatusBadRequest,
			Status: "success",
			Data:   "not found method",
		}
		helper.WriteToResponse(w, response, response.Code)
	}

}

func main() {
	mux := http.NewServeMux()
	category := new(Category)
	mux.Handle("/", category)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
