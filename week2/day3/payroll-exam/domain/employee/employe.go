package employee

import (
	"errors"
	"payroll-exam/db"
	"payroll-exam/domain/salarymatrix"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

type EmployeeRequest struct {
	NameEmployee interface{} `json:"name_employee"`
	Gender       interface{} `json:"gender_employee"`
	Grade        interface{} `json:"grade_employee"`
	IsMarried    interface{} `json:"is_married"`
}

type Employee struct {
	IdEmployee   string `json:"employee_id"`
	NameEmployee string `json:"name_employee"`
	Gender       string `json:"gender"`
	Grade        int    `json:"grade"`
	IsMarried    bool   `json:"is_married"`
}

type EmployeeInterface interface {
	Add(data Employee) (Employee, error)
	ShowAll(channel chan []Employee)
	FindEmplById(emploId string) Employee
}

type ListEmployee struct {
	Employees []Employee `json:"employees"`
}

type EmployeeImplement struct {
	SalaryMatrix salarymatrix.SalaryMatrixInterface
	DB           db.DBInterface
}

func NewEmployee(salaraymatrix salarymatrix.SalaryMatrixInterface, db db.DBInterface) EmployeeInterface {
	return &EmployeeImplement{
		SalaryMatrix: salaraymatrix,
		DB:           db,
	}
}

func (l *EmployeeImplement) Add(data Employee) (Employee, error) {

	salaraymatrix := l.SalaryMatrix.FindByGrade(data.Grade)
	if salaraymatrix == (salarymatrix.SalaryMatrix{}) {
		return Employee{}, errors.New("Grade tidak ditemukan")
	}

	id := "employ-" + uuid.New().String()
	newEmployee := Employee{
		IdEmployee:   id,
		NameEmployee: data.NameEmployee,
		Gender:       data.Gender,
		Grade:        data.Grade,
		IsMarried:    data.IsMarried,
	}

	// type ListEmployee struct {
	// 	Employees []Employee `json:"employees"`
	// }
	listemployee := ListEmployee{}
	l.DB.ReadFile("employee.json", &listemployee)
	listemployee.Employees = append(listemployee.Employees, newEmployee)
	l.DB.MarshalMan("employee.json", listemployee)

	return newEmployee, nil
}

func (l *EmployeeImplement) ShowAll(channel chan []Employee) {
	listEmployee := ListEmployee{}
	l.DB.ReadFile("employee.json", &listEmployee)
	channel <- listEmployee.Employees
}

func (l *EmployeeImplement) FindEmplById(emploId string) Employee {

	listemployee := ListEmployee{}
	l.DB.ReadFile("employee.json", &listemployee)
	for _, v := range listemployee.Employees {
		if v.IdEmployee == emploId {
			return v
		}
	}
	return Employee{}
}

func (er *EmployeeRequest) ValidateFiled() error {

	if er.NameEmployee == nil || er.Gender == nil || er.Grade == nil || er.IsMarried == nil {
		return errors.New("Mohon lengkapi data")
	}
	if !strings.Contains("LP", er.Gender.(string)) || len(er.Gender.(string)) != 1 {
		return errors.New("Mohon isi Gender dengan nilai L/P")
	}

	if reflect.TypeOf(er.NameEmployee).Kind() != reflect.String ||
		reflect.TypeOf(er.Gender).Kind() != reflect.String ||
		reflect.TypeOf(er.IsMarried).Kind() != reflect.Bool ||
		reflect.TypeOf(er.Grade).Kind() != reflect.Float64 {
		return errors.New("Tipe data tidak sesuai")
	}
	return nil
}
