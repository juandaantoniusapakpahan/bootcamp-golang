package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"payroll-exam/domain/employee"
	"payroll-exam/domain/payroll"
	"payroll-exam/domain/salarymatrix"
	"strconv"
)

type JsonResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func GetRequestBody(r *http.Request, data interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func ResponseWrite(w http.ResponseWriter, data interface{}, statuscode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}

}

func EmployeeData(data employee.EmployeeRequest) employee.Employee {
	grade2, _ := strconv.ParseFloat(fmt.Sprintf("%f", data.Grade), 64)
	married1, _ := strconv.ParseBool(fmt.Sprintf("%t", data.IsMarried))

	return employee.Employee{
		NameEmployee: fmt.Sprintf("%s", data.NameEmployee),
		Gender:       fmt.Sprintf("%s", data.Gender),
		Grade:        int(grade2),
		IsMarried:    married1,
	}

}

func SalaryMatrixData(data salarymatrix.SalaryMatrixRequest) salarymatrix.SalaryMatrix {
	grade, _ := strconv.ParseFloat(fmt.Sprintf("%f", data.Grade), 64)
	basicsalary, _ := strconv.ParseFloat(fmt.Sprintf("%f", data.BasicSalary), 64)
	paycut, _ := strconv.ParseFloat(fmt.Sprintf("%f", data.PayCut), 64)
	allowance, _ := strconv.ParseFloat(fmt.Sprintf("%f", data.Allowance), 64)
	headoffamily, _ := strconv.ParseFloat(fmt.Sprintf("%f", data.HeadOfFamily), 64)

	return salarymatrix.SalaryMatrix{
		Grade:        int(grade),
		BasicSalary:  basicsalary,
		PayCut:       paycut,
		Allowance:    allowance,
		HeadOfFamily: headoffamily,
	}
}

func PayrollData(data payroll.PayRollRequest) payroll.PayRollModel {
	hadir, _ := strconv.ParseFloat(fmt.Sprintf("%f", data.Hadir), 64)
	absen, _ := strconv.ParseFloat(fmt.Sprintf("%f", data.Absen), 64)

	return payroll.PayRollModel{
		EmployeeId: fmt.Sprintf("%s", data.EmployeeId),
		Hadir:      hadir,
		Absen:      absen,
	}
}

func NewResponse(code int, status string, data map[string]interface{}) *JsonResponse {
	return &JsonResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
