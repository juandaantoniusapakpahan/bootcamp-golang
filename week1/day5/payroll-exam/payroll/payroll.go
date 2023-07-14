package payroll

import (
	"fmt"
	"payroll-exam/employee"
	"payroll-exam/salarymatrix"
	"strconv"
	"time"
)

// STRUCT
// Payroll → IdPayroll, BasictSalary, PayCut, AdditionalSalary, Employee employee

// Task No 3
// Fitur Tambah Payroll dengan inputan berupa pegawai,
// jumlah hari masuk, jumlah hari tidak masuk yang nanti
// akan diproses menjadi gaji karyawan tersebut

type PayRoll struct {
	IdPayRool        string
	BasicSalary      float64
	PayCut           float64
	AdditionalSalary float64
	Priod            string
	Total            float64
	Employee         employee.Employee
}

type PayRollList struct {
	SalaryMatrix salarymatrix.SalaryMatrixInterface
	Employee     employee.EmployeeInterface
	PayRolls     []PayRoll
}

type PayRollInterface interface {
	Add(employId string, hadir float64, absen float64)
	ShowPayrollById(employId string)
	IsPayroll(emploId string, priod string) bool
}

func NewPayroll(salarymatris salarymatrix.SalaryMatrixInterface, employee employee.EmployeeInterface) PayRollInterface {
	return &PayRollList{SalaryMatrix: salarymatris, Employee: employee}
}

func (p *PayRollList) IsPayroll(emploId string, date string) bool {
	for _, v := range p.PayRolls {
		if v.Priod == date && v.Employee.IdEmployee == emploId {
			return true
		}
	}
	return false
}
func (p *PayRollList) Add(employId string, hadir float64, absen float64) {
	idPayroll := "payroll-" + strconv.Itoa(len(p.PayRolls))
	year, month, _ := time.Now().Date()
	priod := fmt.Sprintf("%v-%v", month, year)

	Dataemploye := p.Employee.FindEmplById(employId)
	if Dataemploye == (employee.Employee{}) {
		fmt.Println("Employee tidak ditemukan")
		return
	}

	salarymatrix := p.SalaryMatrix.FindByGrade(Dataemploye.Grade)

	if isPeriod := p.IsPayroll(Dataemploye.IdEmployee, priod); isPeriod {
		fmt.Println(Dataemploye.NameEmployee, "sudah payroll untuk bulan", month, year)
		return
	}

	var addHeadofFamily float64

	if Dataemploye.Gender == "L" && Dataemploye.IsMarried == true {
		addHeadofFamily = salarymatrix.HeadOfFamily
	}

	newPayroll := PayRoll{
		IdPayRool:        idPayroll,
		BasicSalary:      salarymatrix.BasicSalary,
		PayCut:           absen * salarymatrix.PayCut,
		AdditionalSalary: hadir*salarymatrix.Allowance + addHeadofFamily,
		Priod:            priod,
		Total:            (salarymatrix.BasicSalary - absen*salarymatrix.PayCut + hadir*salarymatrix.Allowance + addHeadofFamily),
		Employee:         Dataemploye,
	}

	p.PayRolls = append(p.PayRolls, newPayroll)
}

func (p *PayRollList) ShowPayrollById(emploId string) {

	profile := 0
	for _, v := range p.PayRolls {

		if v.Employee.IdEmployee == emploId {
			if profile == 0 {
				fmt.Println("Id Employee\t Nama Employe\t Gender\t Grade\t IsMarried\t")
				fmt.Println(
					v.Employee.IdEmployee, "\t",
					v.Employee.NameEmployee, "\t",
					v.Employee.Gender, "\t",
					v.Employee.Grade, "\t",
					v.Employee.IsMarried,
				)
				fmt.Println()
				fmt.Println("Payroll")
				fmt.Println("Id Payroll\t Basic Salary\t PayCut\t AdditionalSalary\t Priod\t\t Total")
				profile += 1
			}

			fmt.Printf("%s\t %.2f\t %.2f\t %.2f\t\t %s\t %.2f\n", v.IdPayRool, v.BasicSalary, v.PayCut, v.AdditionalSalary, v.Priod, v.Total)
		}
	}

}
