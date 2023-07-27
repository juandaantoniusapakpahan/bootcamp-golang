package helper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type JsonResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ReponseError struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

func GetRequestBody(r *http.Request, data interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		panic(err.Error())
	}
	err = json.Unmarshal([]byte(body), data)
	if err != nil {
		panic(err.Error())
	}
}

func ResponseWrite(w http.ResponseWriter, data interface{}, statuscode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
}

func IfPanic(err error) {
	if err != nil {
		panic(err)
	}
}
