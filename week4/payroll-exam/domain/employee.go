package domain

import (
	"fmt"
	"payroll-exam/exception"
	"reflect"
	"strconv"
	"strings"
)

type AddEmployee struct {
	NameEmployee interface{} `json:"name_employee"`
	Gender       interface{} `json:"gender_employee"`
	Grade        interface{} `json:"grade_employee"`
	IsMarried    interface{} `json:"is_married"`
}

type AddedEmployee struct {
	IdEmployee   string `json:"employee_id"`
	NameEmployee string `json:"name_employee"`
	Gender       string `json:"gender"`
	Grade        int    `json:"grade"`
	IsMarried    bool   `json:"is_married"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type EmployeeXPayroll struct {
	IdEmployee   string             `json:"employee_id"`
	NameEmployee string             `json:"name_employee"`
	Gender       string             `json:"gender"`
	Grade        int                `json:"grade"`
	IsMarried    bool               `json:"is_married"`
	CreatedAt    string             `json:"created_at"`
	UpdatedAt    string             `json:"updated_at"`
	Payrolls     []PayRollXEmployee `json:"payrolls"`
}

type Employee struct {
	IdEmployee   string
	NameEmployee string
	Gender       string
	Grade        int
	IsMarried    bool
	CreatedAt    string
	UpdatedAt    string
}

func (er *AddEmployee) ValidateFiled() {

	if er.NameEmployee == nil || er.Gender == nil || er.Grade == nil || er.IsMarried == nil {
		panic(exception.NewBadRequestError("EMPLOYEE.NOT_CONTAIN_NEEDED_PROPERTY"))

	}

	if reflect.TypeOf(er.NameEmployee).Kind() != reflect.String ||
		reflect.TypeOf(er.Gender).Kind() != reflect.String ||
		reflect.TypeOf(er.IsMarried).Kind() != reflect.Bool ||
		reflect.TypeOf(er.Grade).Kind() != reflect.Float64 {
		panic(exception.NewBadRequestError("EMPLOYEE.NOT_MEET_DATA_TYPE_SPECIFICATION"))
	}
	if !strings.Contains("LP", er.Gender.(string)) || len(er.Gender.(string)) != 1 {
		panic(exception.NewBadRequestError("EMPLOYEE.PLEASE_CHOOSE_L_OR_P"))
	}
}

func NewAddedEmployee(employee *Employee) *AddedEmployee {
	return &AddedEmployee{
		IdEmployee:   employee.IdEmployee,
		NameEmployee: employee.NameEmployee,
		Gender:       employee.Gender,
		Grade:        employee.Grade,
		IsMarried:    employee.IsMarried,
		CreatedAt:    employee.CreatedAt,
		UpdatedAt:    employee.UpdatedAt,
	}
}

func (ap *AddEmployee) NewEmployee() *Employee {
	grade, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", ap.Grade), 64)
	gender := fmt.Sprintf("%s", ap.Gender)
	name := fmt.Sprintf("%s", ap.NameEmployee)
	isMarried, _ := strconv.ParseBool(fmt.Sprintf("%t", ap.IsMarried))

	return &Employee{
		NameEmployee: name,
		Gender:       gender,
		Grade:        int(grade),
		IsMarried:    isMarried,
	}
}
