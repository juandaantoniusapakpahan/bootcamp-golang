package exception

import (
	"encoding/json"
	"net/http"
	"payroll-exam/helper"
)

func ErrorHandling(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if ser, ok := err.(BadRequestError); ok {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				response := helper.ReponseError{
					Code:    http.StatusBadRequest,
					Status:  "fail",
					Message: ser.Error,
				}
				encoder := json.NewEncoder(w)
				encoder.Encode(response)

			} else if ser, ok := err.(NotFoundError); ok {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotFound)
				response := helper.ReponseError{
					Code:    http.StatusNotFound,
					Status:  "fail",
					Message: ser.Error,
				}
				encoder := json.NewEncoder(w)
				encoder.Encode(response)
			} else if err != nil {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				response := helper.ReponseError{
					Code:    http.StatusInternalServerError,
					Status:  "error",
					Message: "Maaf, terjadi kesalah pada server",
				}
				encoder := json.NewEncoder(w)
				encoder.Encode(response)

			}
		}()
		next.ServeHTTP(w, r)
	})

}
