package helper

import (
	"encoding/json"
	"net/http"
)

type ResponseTemplate struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func GetResponseBody(r *http.Request, data interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		panic(err)
	}
}
