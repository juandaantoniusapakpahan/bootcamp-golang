package domain

import (
	"fmt"
	"payroll-exam/exception"
	"reflect"
	"strconv"
)

type AddPayRoll struct {
	EmployeeId interface{} `json:"employee_id"`
	Priode     interface{} `json:"priode"`
	Hadir      interface{} `json:"hadir"`
	Absen      interface{} `json:"absen"`
}

type AddedPayRoll struct {
	IdPayRool        string  `json:"payroll_id"`
	Priode           string  `json:"priode"`
	JumlahHadir      int     `json:"jumlah_hadir"`
	JumlahAbsen      int     `json:"jumlah_absen"`
	BasicSalary      float64 `json:"basic_salary"`
	PayCut           float64 `json:"pay_cut"`
	AdditionalSalary float64 `json:"additional_salary"`
	HeadOfFamily     float64 `json:"head_of_familly"`
	Total            float64 `json:"total"`
	EmployeeId       string  `json:"employee_id"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}
type PayRollXEmployee struct {
	IdPayRool        string  `json:"payroll_id"`
	Priode           string  `json:"priode"`
	JumlahHadir      int     `json:"jumlah_hadir"`
	JumlahAbsen      int     `json:"jumlah_absen"`
	BasicSalary      float64 `json:"basic_salary"`
	PayCut           float64 `json:"pay_cut"`
	AdditionalSalary float64 `json:"additional_salary"`
	HeadOfFamily     float64 `json:"head_of_familly"`
	Total            float64 `json:"total"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

type PayRoll struct {
	IdPayRool        string
	Priode           string
	JumlahHadir      int
	JumlahAbsen      int
	BasicSalary      float64
	PayCut           float64
	AdditionalSalary float64
	HeadOfFamily     float64
	Total            float64
	EmployeeId       string
	CreatedAt        string
	UpdatedAt        string
}

type GetPayrollById struct {
	AddedEmployee
	Payroll []*PayRollXEmployee `json:"payrolls"`
}

func (pr *AddPayRoll) ValidateFiled() {
	if pr.EmployeeId == nil || pr.Priode == nil || pr.Hadir == nil || pr.Absen == nil {
		panic(exception.NewBadRequestError("PAYROLL.NOT_CONTAIN_NEEDED_PROPERTY"))
	}
	if reflect.TypeOf(pr.EmployeeId).Kind() != reflect.String ||
		reflect.TypeOf(pr.Priode).Kind() != reflect.String ||
		reflect.TypeOf(pr.Hadir).Kind() != reflect.Float64 ||
		reflect.TypeOf(pr.Absen).Kind() != reflect.Float64 {
		panic(exception.NewBadRequestError("PAYROLL.NOT_MEET_DATA_TYPE_SPECIFICATION"))
	}

}

func (pr *AddPayRoll) NewPayroll() *PayRoll {
	employeeId := fmt.Sprintf("%s", pr.EmployeeId)
	priode := fmt.Sprintf("%s", pr.Priode)
	hadir, _ := strconv.ParseFloat(fmt.Sprintf("%.f", pr.Hadir), 64)
	absen, _ := strconv.ParseFloat(fmt.Sprintf("%.f", pr.Absen), 64)

	return &PayRoll{
		EmployeeId:  employeeId,
		Priode:      priode,
		JumlahHadir: int(hadir),
		JumlahAbsen: int(absen),
	}
}

func NewAddedPayroll(py *PayRoll) *AddedPayRoll {
	return &AddedPayRoll{
		IdPayRool:        py.IdPayRool,
		Priode:           py.Priode,
		JumlahHadir:      py.JumlahHadir,
		JumlahAbsen:      py.JumlahAbsen,
		BasicSalary:      py.BasicSalary,
		PayCut:           py.PayCut,
		AdditionalSalary: py.AdditionalSalary,
		HeadOfFamily:     py.HeadOfFamily,
		Total:            py.Total,
		EmployeeId:       py.EmployeeId,
		CreatedAt:        py.CreatedAt,
		UpdatedAt:        py.UpdatedAt,
	}
}
