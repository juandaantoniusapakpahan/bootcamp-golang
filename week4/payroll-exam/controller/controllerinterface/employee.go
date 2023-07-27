package controllerinterface

import "net/http"

type EmployeeControllerInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
