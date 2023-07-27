package routes

import (
	"payroll-exam/controller/controllerinterface"
	"payroll-exam/exception"
	"payroll-exam/middleware"
)

func NewRoute(
	csm controllerinterface.SalaryMatrixControllerInterface,
	cep controllerinterface.EmployeeControllerInterface,
	pc controllerinterface.PayrollControllerInterface,
) *middleware.CustomMux {
	mux := new(middleware.CustomMux)

	mux.Handle("/salarymatrix", csm)
	mux.Handle("/employee", cep)
	mux.Handle("/payroll", pc)

	mux.RegisterMiddleware(middleware.SetMyCookies)
	mux.RegisterMiddleware(exception.ErrorHandling)
	// mux.RegisterMiddleware(middleware.LogRequest)

	return mux
}
