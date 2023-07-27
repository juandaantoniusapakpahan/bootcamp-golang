package controllerinterface

import "net/http"

type PayrollControllerInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
