package main

import (
	"fmt"
	"log"
	"net/http"
	"payroll-exam/db"
	"payroll-exam/domain/employee"
	"payroll-exam/domain/payroll"
	"payroll-exam/domain/salarymatrix"
	"payroll-exam/handler"
	"payroll-exam/routes"
	"runtime"
)

func main() {
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("THREAD:", runtime.GOMAXPROCS(-1))
	fmt.Println("GOROUTINE:", runtime.NumGoroutine())
	// Go di ShowAll Employee

	newDB := db.NewDB()
	salaryMatrix := salarymatrix.NewSalaryMatrix(newDB)
	employ := employee.NewEmployee(salaryMatrix, newDB)
	newPayroll := payroll.NewPayroll(newDB, salaryMatrix, employ)

	// handler
	employeHandler := handler.NewEmployeServe(employ)
	salaryMatrixHandler := handler.NewSelaryMatrixHandler(salaryMatrix)
	payrollHandler := handler.NewPayrollHandler(newPayroll)
	userHandler := handler.NewUserHandler()

	newDB.OpenLogFile("db/development.log")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// route
	route := routes.NewRoute(employeHandler, salaryMatrixHandler, payrollHandler, userHandler)

	// server
	server := http.Server{
		Addr:    ":8080",
		Handler: route,
	}

	// run server
	fmt.Println("Server started at localhost:8080")
	server.ListenAndServe()
	fmt.Println(runtime.NumGoroutine())

}
