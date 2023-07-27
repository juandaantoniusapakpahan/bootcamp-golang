package handler

import (
	"net/http"
	"payroll-exam/domain/salarymatrix"
	"payroll-exam/exception"
	"payroll-exam/helper"
)

type SalaryMatrixHandler struct {
	SalaryMatrix salarymatrix.SalaryMatrixInterface
}

type SalaryMatrixHandlerInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func NewSelaryMatrixHandler(salarymatrix salarymatrix.SalaryMatrixInterface) SalaryMatrixHandlerInterface {
	return &SalaryMatrixHandler{
		SalaryMatrix: salarymatrix,
	}
}

func (sh *SalaryMatrixHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	grade := r.URL.Query().Get("grade")

	if r.Method == "POST" && grade == "" {
		requestBody := salarymatrix.SalaryMatrixRequest{}
		helper.GetRequestBody(r, &requestBody)

		err := requestBody.ValidateFiled()
		if err != nil {
			panic(exception.NewBadRequestError(err.Error()))
		}

		salaryMatrixData := helper.SalaryMatrixData(requestBody)

		result, err := sh.SalaryMatrix.Add(salaryMatrixData)
		if err != nil {
			panic(exception.NewBadRequestError(err.Error()))
		}

		response := helper.NewResponse(http.StatusCreated, "success", map[string]interface{}{"salary_matrix": result})

		helper.ResponseWrite(w, response, http.StatusCreated)
	} else if r.Method == "GET" {

		channel := make(chan []salarymatrix.SalaryMatrix)
		go sh.SalaryMatrix.GetAll(channel)
		result := <-channel
		close(channel)

		response := helper.NewResponse(http.StatusOK, "success", map[string]interface{}{"salary_matrix": result})

		helper.ResponseWrite(w, response, http.StatusOK)
	} else {
		response := helper.NewResponse(http.StatusBadRequest, "fail", map[string]interface{}{"message": "method not allowed"})

		helper.ResponseWrite(w, response, http.StatusBadRequest)
	}
}
