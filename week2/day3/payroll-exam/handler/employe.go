package handler

import (
	"net/http"
	"payroll-exam/domain/employee"
	"payroll-exam/exception"
	"payroll-exam/helper"
)

type EmployeServeHTTP struct {
	Employee employee.EmployeeInterface
}

type EmployeServeHTTPInter interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type HandlerEmployee struct {
	Employee employee.EmployeeInterface
}

func NewEmployeServe(employ employee.EmployeeInterface) EmployeServeHTTPInter {
	return &EmployeServeHTTP{
		Employee: employ,
	}
}

func (he *EmployeServeHTTP) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	employId := r.URL.Query().Get("employId")

	if r.Method == "GET" && employId == "" {

		employeChan := make(chan []employee.Employee)
		go he.Employee.ShowAll(employeChan)
		employees := <-employeChan
		close(employeChan)

		response := helper.NewResponse(http.StatusOK, "success", map[string]interface{}{"employees": employees})
		helper.ResponseWrite(w, response, http.StatusOK)
	} else if r.Method == "POST" {
		requestBody := employee.EmployeeRequest{}
		helper.GetRequestBody(r, &requestBody)

		err := requestBody.ValidateFiled()
		if err != nil {
			panic(exception.NewBadRequestError(err.Error()))
		}

		employeeData := helper.EmployeeData(requestBody) // from interface to stuct
		result, err := he.Employee.Add(employeeData)

		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		response := helper.NewResponse(http.StatusCreated, "success", map[string]interface{}{"employee": result})
		helper.ResponseWrite(w, response, http.StatusCreated)
	} else if r.Method == "GET" && employId != "" {
		result := he.Employee.FindEmplById(employId)
		response := helper.NewResponse(http.StatusOK, "success", map[string]interface{}{"employee": result})
		helper.ResponseWrite(w, response, http.StatusOK)
	} else {
		response := helper.NewResponse(http.StatusBadRequest, "fail", map[string]interface{}{"message": "Method not allowed"})
		helper.ResponseWrite(w, response, http.StatusBadRequest)
	}
}
