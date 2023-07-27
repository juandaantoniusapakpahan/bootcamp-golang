package controller

import (
	"net/http"
	"payroll-exam/controller/controllerinterface"
	"payroll-exam/domain"
	"payroll-exam/helper"
	"payroll-exam/service/serviceinterface"
)

type EmployeeControllerImplement struct {
	EmployeeService serviceinterface.EmployeeServiceInterface
}

func NewEmployeeControllerImplement(employeeService serviceinterface.EmployeeServiceInterface) controllerinterface.EmployeeControllerInterface {
	return &EmployeeControllerImplement{
		EmployeeService: employeeService,
	}
}

func (ec *EmployeeControllerImplement) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		requestBody := domain.AddEmployee{}
		helper.GetRequestBody(r, &requestBody)

		result := ec.EmployeeService.Create(r.Context(), &requestBody)

		response := helper.JsonResponse{
			Code:   http.StatusCreated,
			Status: "success",
			Data:   map[string]interface{}{"employee": result},
		}

		helper.ResponseWrite(w, response, response.Code)
	case "GET":
		employeeId := r.URL.Query().Get("employeeId")
		switch {
		case employeeId != "":
			employeeId := r.URL.Query().Get("employeeId")

			result := ec.EmployeeService.FindById(r.Context(), employeeId)

			response := helper.JsonResponse{
				Code:   http.StatusCreated,
				Status: "success",
				Data:   map[string]interface{}{"employee": result},
			}
			helper.ResponseWrite(w, response, response.Code)

		default:
			result := ec.EmployeeService.FindAll(r.Context())
			response := helper.JsonResponse{
				Code:   http.StatusCreated,
				Status: "success",
				Data:   map[string]interface{}{"employees": result},
			}
			helper.ResponseWrite(w, response, response.Code)

		}

	}
}
