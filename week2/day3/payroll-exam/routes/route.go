package routes

import (
	"payroll-exam/exception"
	"payroll-exam/handler"
	"payroll-exam/middleware"
)

func NewRoute(emplyeserve handler.EmployeServeHTTPInter,
	salaryMatrix handler.SalaryMatrixHandlerInterface,
	payrollHandler handler.PayrollHandlerInterface,
	userHandler handler.UserHandlerInterface,
) *middleware.CustomMux {
	mux := new(middleware.CustomMux)

	mux.Handle("/employee", emplyeserve)
	mux.Handle("/salarymatrix", salaryMatrix)
	mux.Handle("/payroll", payrollHandler)
	mux.Handle("/cookies", userHandler)

	// var handler http.Handler = mux
	// handler = middleware.SetMyCookies(handler)

	mux.RegisterMiddleware(middleware.SetMyCookies)
	mux.RegisterMiddleware(exception.ErrorHandling)
	mux.RegisterMiddleware(middleware.LogRequest)

	return mux
}
