package handler

import (
	"net/http"
	"payroll-exam/helper"
	"time"
)

type UserHandler struct {
}

type UserHandlerInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler() UserHandlerInterface {
	return &UserHandler{}
}

func (uh *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		c := &http.Cookie{}
		c.Name = "PayRoll"
		c.Expires = time.Unix(0, 0)
		c.MaxAge = -1
		http.SetCookie(w, c)
		response := helper.JsonResponse{
			Code:   http.StatusCreated,
			Status: "success",
			Data:   map[string]interface{}{"message": "Berhasil logout"},
		}
		helper.ResponseWrite(w, response, http.StatusCreated)
	} else {
		response := helper.JsonResponse{
			Code:   http.StatusBadRequest,
			Status: "success",
			Data:   map[string]interface{}{"message": "method not allowed"},
		}
		helper.ResponseWrite(w, response, http.StatusBadRequest)
	}
}
