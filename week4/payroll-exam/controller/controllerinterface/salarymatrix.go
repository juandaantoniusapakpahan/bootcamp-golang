package controllerinterface

import (
	"net/http"
)

type SalaryMatrixControllerInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
