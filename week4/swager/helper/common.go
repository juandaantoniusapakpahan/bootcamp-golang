package helper

import (
	"encoding/json"
	"net/http"
)

type ResponseTemplete struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func GetRequestBody(r *http.Request, data interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		panic(err)
	}
}

func WriteToResponse(w http.ResponseWriter, data interface{}, statuscode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}
}
