package routes

import (
	"net/http"
	"payroll-exam/exception"
	"payroll-exam/handler"
	"payroll-exam/middleware"
)

type CustomMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		current = next(current)
	}
	current.ServeHTTP(w, r)
}

func NewRoute(emplyeserve handler.EmployeServeHTTPInter,
	salaryMatrix handler.SalaryMatrixHandlerInterface,
	payrollHandler handler.PayrollHandlerInterface,
	userHandler handler.UserHandlerInterface,
) *CustomMux {
	mux := new(CustomMux)

	mux.Handle("/employee", emplyeserve)
	mux.Handle("/salarymatrix", salaryMatrix)
	mux.Handle("/payroll", payrollHandler)
	mux.Handle("/cookies", userHandler)

	// var handler http.Handler = mux
	// handler = middleware.SetMyCookies(handler)

	mux.RegisterMiddleware(middleware.SetMyCookies)
	mux.RegisterMiddleware(exception.ErrorHandling)

	return mux
}
