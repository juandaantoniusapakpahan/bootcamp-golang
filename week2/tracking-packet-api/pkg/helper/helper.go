package helper

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func PanicHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func GetRequestBody(r *http.Request, data interface{}) error {
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&data)
	return err
}

func ChangeToJson(w http.ResponseWriter, data interface{}) {
	encode := json.NewEncoder(w)
	err := encode.Encode(data)
	PanicHandler(err)
}

func RollBackCommit(tx *sql.Tx) {
	err := recover()

	if err != nil {
		errRol := tx.Rollback()
		PanicHandler(errRol)
		panic(err)

	} else {
		err := tx.Commit()
		PanicHandler(err)
	}
}
