package controller

import (
	"net/http"
	"payroll-exam/controller/controllerinterface"
	"payroll-exam/domain"
	"payroll-exam/helper"
	"payroll-exam/service/serviceinterface"
)

type SalaryMatrixControllerImplement struct {
	SalaryMatrixService serviceinterface.SalaryMatrixServiceInterface
}

func NewSalaryMatrixControllerImplement(service serviceinterface.SalaryMatrixServiceInterface) controllerinterface.SalaryMatrixControllerInterface {
	return &SalaryMatrixControllerImplement{SalaryMatrixService: service}
}

func (sc *SalaryMatrixControllerImplement) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		salarymatrixId := r.URL.Query().Get("salarymatrixid")
		switch {
		case salarymatrixId != "":
			result := sc.SalaryMatrixService.FindById(r.Context(), salarymatrixId)

			response := helper.JsonResponse{
				Code:   http.StatusOK,
				Status: "success",
				Data:   map[string]interface{}{"salarymatrix": result},
			}
			helper.ResponseWrite(w, response, http.StatusOK)
		case salarymatrixId == "":
			result := sc.SalaryMatrixService.FindAll(r.Context())
			response := helper.JsonResponse{
				Code:   http.StatusOK,
				Status: "success",
				Data:   map[string]interface{}{"salarymatrixs": result},
			}
			helper.ResponseWrite(w, response, http.StatusOK)
		}

	case "POST":
		requestBody := domain.AddSalaryMatrix{}
		helper.GetRequestBody(r, &requestBody)
		result := sc.SalaryMatrixService.Create(r.Context(), requestBody)
		response := helper.JsonResponse{
			Code:   http.StatusCreated,
			Status: "success",
			Data:   map[string]interface{}{"salarymatrix": result},
		}
		helper.ResponseWrite(w, response, response.Code)

	case "PUT":
		salarymatrixId := r.URL.Query().Get("salarymatrixid")
		requestBody := domain.AddSalaryMatrix{}
		helper.GetRequestBody(r, &requestBody)

		result := sc.SalaryMatrixService.Edit(r.Context(), requestBody, salarymatrixId)
		response := helper.JsonResponse{
			Code:   http.StatusOK,
			Status: "success",
			Data:   map[string]interface{}{"salarymatrix": result},
		}
		helper.ResponseWrite(w, response, response.Code)
	}
}
