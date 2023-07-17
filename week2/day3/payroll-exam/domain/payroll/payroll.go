package payroll

import (
	"errors"
	"os"
	"payroll-exam/db"
	"payroll-exam/domain/employee"
	"payroll-exam/domain/salarymatrix"
	"reflect"
	"sync"

	"github.com/google/uuid"
)

type PayRollRequest struct {
	EmployeeId interface{} `json:"employee_id"`
	Priode     interface{} `json:"priode"`
	Hadir      interface{} `json:"hadir"`
	Absen      interface{} `json:"absen"`
}

func (pr *PayRollRequest) ValidateFiled() error {
	if pr.EmployeeId == nil || pr.Priode == nil || pr.Hadir == nil || pr.Absen == nil {
		return errors.New("Mohon lengkapi data")
	}
	if reflect.TypeOf(pr.EmployeeId).Kind() != reflect.String ||
		reflect.TypeOf(pr.Priode).Kind() != reflect.String ||
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
	Priode           string            `json:"priode"`
	Total            float64           `json:"total"`
	Employee         employee.Employee `json:"employee"`
}

type PayRollModel struct {
	EmployeeId string
	Hadir      float64
	Absen      float64
	Priode     string
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
	ShowPayrollById(employId string) (ShowPayRollById, error)
	IsPayroll(emploId string, priod string) bool
	Download() *os.File
	ShowAll(channel chan []PayRoll, mg *sync.WaitGroup)
}

func NewPayroll(db db.DBInterface, salarymatris salarymatrix.SalaryMatrixInterface, employee employee.EmployeeInterface) PayRollInterface {
	return &PayrollImplement{DB: db, SalaryMatrix: salarymatris, Employee: employee}
}

func (p *PayrollImplement) IsPayroll(emploId string, date string) bool {
	listPayroll := PayRollList{}
	p.DB.ReadFile("payroll.json", &listPayroll)
	for _, v := range listPayroll.PayRolls {
		if v.Priode == date && v.Employee.IdEmployee == emploId {
			return true
		}
	}
	return false
}

func (p *PayrollImplement) Add(pr PayRollModel) (PayRoll, error) {
	idPayroll := "payroll-" + uuid.New().String()

	Dataemploye := p.Employee.FindEmplById(pr.EmployeeId)
	if Dataemploye == (employee.Employee{}) {
		return PayRoll{}, errors.New("Employe tidak ditemukan")
	}

	salarymatrix := p.SalaryMatrix.FindByGrade(Dataemploye.Grade)

	if isPeriod := p.IsPayroll(Dataemploye.IdEmployee, pr.Priode); isPeriod {
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
		Priode:           pr.Priode,
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

type JustPayRoll struct {
	IdPayRool        string  `json:"payroll_id"`
	BasicSalary      float64 `json:"basic_salary"`
	PayCut           float64 `json:"pay_cut"`
	AdditionalSalary float64 `json:"additional_salary"`
	Priode           string  `json:"priode"`
	Total            float64 `json:"total"`
}

type ShowPayRollById struct {
	Employee employee.Employee `json:"employee"`
	PayRoll  []JustPayRoll     `json:"payrolls"`
}

func (p *PayrollImplement) ShowPayrollById(emploId string) (ShowPayRollById, error) {
	listPayroll := PayRollList{}
	payrollRespon := ShowPayRollById{}
	p.DB.ReadFile("payroll.json", &listPayroll)

	dataEmployee := p.Employee.FindEmplById(emploId)
	if dataEmployee == (employee.Employee{}) {
		return ShowPayRollById{}, errors.New("Payroll tidak ditemukan")
	}

	payrollRespon.Employee = dataEmployee

	for _, v := range listPayroll.PayRolls {
		if v.Employee.IdEmployee == emploId {
			newPayjus := JustPayRoll{
				IdPayRool:        v.IdPayRool,
				BasicSalary:      v.BasicSalary,
				PayCut:           v.PayCut,
				AdditionalSalary: v.AdditionalSalary,
				Priode:           v.Priode,
				Total:            v.Total,
			}
			payrollRespon.PayRoll = append(payrollRespon.PayRoll, newPayjus)
		}
	}
	return payrollRespon, nil
}

func (p *PayrollImplement) ShowAll(channel chan []PayRoll, mg *sync.WaitGroup) {
	mg.Add(1)
	listPayroll := PayRollList{}
	p.DB.ReadFile("payroll.json", &listPayroll)
	channel <- listPayroll.PayRolls
	mg.Done()

}
