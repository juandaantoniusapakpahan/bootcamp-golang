package handler

import (
	"fmt"
	"io"
	"net/http"
	"payroll-exam/domain/payroll"
	"payroll-exam/exception"
	"payroll-exam/helper"
)

type PayrollHandler struct {
	Payroll payroll.PayRollInterface
}

type PayrollHandlerInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func NewPayrollHandler(payroll payroll.PayRollInterface,
) PayrollHandlerInterface {
	return &PayrollHandler{
		Payroll: payroll,
	}
}

func (ph *PayrollHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	employeId := r.URL.Query().Get("employeeId")
	download := r.URL.Query().Get("download")

	if r.Method == "POST" && employeId == "" {

		bodyRequest := payroll.PayRollRequest{}
		helper.GetRequestBody(r, &bodyRequest)

		err := bodyRequest.ValidateFiled()
		if err != nil {

			panic(exception.NewBadRequestError(err.Error()))
		}

		payrollData := helper.PayrollData(bodyRequest)
		result, err := ph.Payroll.Add(payrollData)

		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}

		response := helper.NewResponse(http.StatusCreated, "success", map[string]interface{}{"payroll": result})
		helper.ResponseWrite(w, response, http.StatusCreated)

	} else if r.Method == "GET" && employeId != "" {

		result, err := ph.Payroll.ShowPayrollById(employeId)
		if err != nil {
			panic(exception.NewNotFoundError(err.Error()))
		}
		response := helper.NewResponse(http.StatusOK, "success", map[string]interface{}{"payroll": result})
		helper.ResponseWrite(w, response, http.StatusOK)

	} else if r.Method == "GET" && download == "" && employeId == "" {
		result := ph.Payroll.ShowAll()
		response := helper.NewResponse(http.StatusCreated, "success", map[string]interface{}{"payrolls": result})
		helper.ResponseWrite(w, response, http.StatusOK)

	} else if r.Method == "GET" && download == "true" {
		file := ph.Payroll.Download()
		w.Header().Set("Content-Disposition", "attachment; filename="+file.Name())

		_, err := io.Copy(w, file)
		if err != nil {
			panic(err)
		}
		fmt.Fprint(w, "GGWP")

	} else {
		response := helper.NewResponse(http.StatusBadRequest, "fail", map[string]interface{}{"message": "Method not allowed"})
		helper.ResponseWrite(w, response, http.StatusBadRequest)
	}

}
