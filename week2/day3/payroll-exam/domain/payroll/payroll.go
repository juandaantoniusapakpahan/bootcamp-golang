package payroll

import (
	"errors"
	"fmt"
	"os"
	"payroll-exam/db"
	"payroll-exam/domain/employee"
	"payroll-exam/domain/salarymatrix"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type PayRollRequest struct {
	EmployeeId interface{} `json:"employee_id"`
	Hadir      interface{} `json:"hadir"`
	Absen      interface{} `json:"absen"`
}

func (pr *PayRollRequest) ValidateFiled() error {
	if pr.EmployeeId == nil || pr.Hadir == nil || pr.Absen == nil {
		return errors.New("Mohon lengkapi data")
	}
	if reflect.TypeOf(pr.EmployeeId).Kind() != reflect.String ||
		reflect.TypeOf(pr.Hadir).Kind() != reflect.Float64 ||
		reflect.TypeOf(pr.Absen).Kind() != reflect.Float64 {
		return errors.New("Tipe data tidak sesuai")
	}

	return nil
}

type PayRoll struct {
	IdPayRool        string            `json:"payroll_id"`
	BasicSalary      float64           `json:"basic_salary"`
	PayCut           float64           `json:"pay_cut"`
	AdditionalSalary float64           `json:"additional_salary"`
	Priod            string            `json:"priod"`
	Total            float64           `json:"total"`
	Employee         employee.Employee `json:"employee"`
}

type PayRollModel struct {
	EmployeeId string
	Hadir      float64
	Absen      float64
}

type PayRollList struct {
	PayRolls []PayRoll `json:"payrolls"`
}
type PayrollImplement struct {
	DB           db.DBInterface
	SalaryMatrix salarymatrix.SalaryMatrixInterface
	Employee     employee.EmployeeInterface
}

type PayRollInterface interface {
	Add(payroll PayRollModel) (PayRoll, error)
	ShowPayrollById(employId string) (PayRoll, error)
	IsPayroll(emploId string, priod string) bool
	Download() *os.File
	ShowAll() []PayRoll
}

func NewPayroll(db db.DBInterface, salarymatris salarymatrix.SalaryMatrixInterface, employee employee.EmployeeInterface) PayRollInterface {
	return &PayrollImplement{DB: db, SalaryMatrix: salarymatris, Employee: employee}
}

func (p *PayrollImplement) IsPayroll(emploId string, date string) bool {
	listPayroll := PayRollList{}
	p.DB.ReadFile("payroll.json", &listPayroll)
	for _, v := range listPayroll.PayRolls {
		if v.Priod == date && v.Employee.IdEmployee == emploId {
			return true
		}
	}
	return false
}

func (p *PayrollImplement) Add(pr PayRollModel) (PayRoll, error) {
	idPayroll := "payroll-" + uuid.New().String()
	year, month, _ := time.Now().Date()
	priod := fmt.Sprintf("%v-%v", month, year)

	Dataemploye := p.Employee.FindEmplById(pr.EmployeeId)
	if Dataemploye == (employee.Employee{}) {
		return PayRoll{}, errors.New("Employe tidak ditemukan")
	}

	salarymatrix := p.SalaryMatrix.FindByGrade(Dataemploye.Grade)

	if isPeriod := p.IsPayroll(Dataemploye.IdEmployee, priod); isPeriod {
		return PayRoll{}, errors.New("Payroll priode sudah dimasukan")
	}

	var addHeadofFamily float64

	if Dataemploye.Gender == "L" && Dataemploye.IsMarried == true {
		addHeadofFamily = salarymatrix.HeadOfFamily
	}

	newPayroll := PayRoll{
		IdPayRool:        idPayroll,
		BasicSalary:      salarymatrix.BasicSalary,
		PayCut:           pr.Absen * salarymatrix.PayCut,
		AdditionalSalary: pr.Hadir*salarymatrix.Allowance + addHeadofFamily,
		Priod:            priod,
		Total:            (salarymatrix.BasicSalary - pr.Absen*salarymatrix.PayCut + pr.Hadir*salarymatrix.Allowance + addHeadofFamily),
		Employee:         Dataemploye,
	}

	listPayroll := PayRollList{}

	p.DB.ReadFile("payroll.json", &listPayroll)
	listPayroll.PayRolls = append(listPayroll.PayRolls, newPayroll)
	p.DB.MarshalMan("payroll.json", listPayroll)

	return newPayroll, nil
}

func (p *PayrollImplement) Download() *os.File {
	file := p.DB.GetFile("payroll.json")
	return file
}

func (p *PayrollImplement) ShowPayrollById(emploId string) (PayRoll, error) {
	listPayroll := PayRollList{}
	p.DB.ReadFile("payroll.json", &listPayroll)
	for _, v := range listPayroll.PayRolls {
		if v.Employee.IdEmployee == emploId {
			return v, nil
		}
	}
	return PayRoll{}, errors.New("Payroll tidak ditemukan")
}

func (p *PayrollImplement) ShowAll() []PayRoll {
	listPayroll := PayRollList{}
	p.DB.ReadFile("payroll.json", &listPayroll)
	return listPayroll.PayRolls
}
