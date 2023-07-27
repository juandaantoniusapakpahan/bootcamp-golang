package controller

import (
	"net/http"
	"payroll-exam/controller/controllerinterface"
	"payroll-exam/domain"
	"payroll-exam/helper"
	"payroll-exam/service/serviceinterface"
)

type PayrollControllerImplement struct {
	PayrollService serviceinterface.PayrollServiceInterface
}

func NewPayrollControllerImplement(ps serviceinterface.PayrollServiceInterface) controllerinterface.PayrollControllerInterface {
	return &PayrollControllerImplement{
		PayrollService: ps,
	}
}

func (pc *PayrollControllerImplement) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":

		requestBody := domain.AddPayRoll{}

		helper.GetRequestBody(r, &requestBody)

		result := pc.PayrollService.Create(r.Context(), &requestBody)
		response := helper.JsonResponse{
			Code:   http.StatusCreated,
			Status: "success",
			Data:   map[string]interface{}{"payroll": result},
		}
		helper.ResponseWrite(w, response, response.Code)
	case "GET":
		employeeId := r.URL.Query().Get("employeeId")
		payrollId := r.URL.Query().Get("payrollId")
		switch {
		case employeeId != "" && payrollId == "":
			result := pc.PayrollService.FindByEmployeeId(r.Context(), employeeId)

			response := helper.JsonResponse{
				Code:   http.StatusOK,
				Status: "success",
				Data:   map[string]interface{}{"payroll": result},
			}
			helper.ResponseWrite(w, response, response.Code)
		case employeeId == "" && payrollId == "":
			result := pc.PayrollService.FindAll(r.Context())

			response := helper.JsonResponse{
				Code:   http.StatusOK,
				Status: "success",
				Data:   map[string]interface{}{"payrolls": result},
			}
			helper.ResponseWrite(w, response, response.Code)

		case employeeId == "" && payrollId != "":

			result := pc.PayrollService.FindById(r.Context(), payrollId)

			response := helper.JsonResponse{
				Code:   http.StatusOK,
				Status: "success",
				Data:   map[string]interface{}{"payrolls": result},
			}
			helper.ResponseWrite(w, response, response.Code)
		}

	}
}
