package main

import (
	"fmt"
	"net/http"
	"os"
	"payroll-exam/controller"
	"payroll-exam/db"
	"payroll-exam/helper"
	"payroll-exam/repository"
	"payroll-exam/routes"
	"payroll-exam/service"

	"runtime"
)

func main() {

	// port := os.Getenv("PORT")

	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("THREAD:", runtime.GOMAXPROCS(-1))
	fmt.Println("GOROUTINE:", runtime.NumGoroutine())

	MyDB := db.ConnectDB()
	newSalaryMatrixRepo := repository.NewSalaryMatrixImplement()
	newSalaryMatrixService := service.NewSalaryMatrixServiceImplement(newSalaryMatrixRepo, MyDB)
	newSalaryMatrixController := controller.NewSalaryMatrixControllerImplement(newSalaryMatrixService)

	newEmployeeRepo := repository.NewEmployeeRepositoryImplement()
	newEmployeeService := service.NewEmployeeServiceImplement(newEmployeeRepo, MyDB)
	newEmployeeController := controller.NewEmployeeControllerImplement(newEmployeeService)

	newPayrollRepo := repository.NewPayrollRepositoryImpolement()
	newPayrollService := service.NewPayrollServiceImplement(MyDB, newPayrollRepo, newEmployeeRepo, newSalaryMatrixRepo)
	newPayrollController := controller.NewPayrollControllerImplement(newPayrollService)

	helper.LogFile("logging.log")

	// route
	route := routes.NewRoute(newSalaryMatrixController, newEmployeeController, newPayrollController)

	http.ListenAndServeTLS(":443", "server.crt", "server.key", helper.Logging(os.Stderr, route))

	// server
	fmt.Println("Server started at localhost:443")

}
