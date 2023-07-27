package domain

import (
	"fmt"
	"payroll-exam/exception"
	"reflect"
	"strconv"
)

type SalaryMatrix struct {
	IdSalary     string
	Grade        int
	BasicSalary  float64
	PayCut       float64
	Allowance    float64
	HeadOfFamily float64
	CreatedAt    string
	UpdatedAt    string
}

type AddSalaryMatrix struct {
	Grade        any `json:"grade"`
	BasicSalary  any `json:"basic_salary"`
	PayCut       any `json:"pay_cut"`
	Allowance    any `json:"allowance"`
	HeadOfFamily any `json:"head_of_family"`
}

type AddedSalaryMatrix struct {
	IdSalary     string  `json:"salary_matrix_id"`
	Grade        int     `json:"grade"`
	BasicSalary  float64 `json:"basic_salary"`
	PayCut       float64 `json:"pay_cut"`
	Allowance    float64 `json:"allowance"`
	HeadOfFamily float64 `json:"head_of_family"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
type UpdatedSalaryMatrix struct {
	IdSalary     string  `json:"salary_matrix_id"`
	Grade        int     `json:"grade"`
	BasicSalary  float64 `json:"basic_salary"`
	PayCut       float64 `json:"pay_cut"`
	Allowance    float64 `json:"allowance"`
	HeadOfFamily float64 `json:"head_of_family"`
	UpdatedAt    string  `json:"updated_at"`
}
type SalaryMatrixList struct {
	SalaryMatrixs []SalaryMatrix `json:"salary_matrix"`
}

func NewAddedSalaryMatrix(salarymatrix *SalaryMatrix) *AddedSalaryMatrix {
	return &AddedSalaryMatrix{
		IdSalary:     salarymatrix.IdSalary,
		Grade:        salarymatrix.Grade,
		BasicSalary:  salarymatrix.BasicSalary,
		PayCut:       salarymatrix.PayCut,
		Allowance:    salarymatrix.Allowance,
		HeadOfFamily: salarymatrix.HeadOfFamily,
		CreatedAt:    salarymatrix.CreatedAt,
		UpdatedAt:    salarymatrix.UpdatedAt,
	}
}

func NewUpdatedSalaryMatrix(salarymatrix *SalaryMatrix) *UpdatedSalaryMatrix {
	return &UpdatedSalaryMatrix{
		IdSalary:     salarymatrix.IdSalary,
		Grade:        salarymatrix.Grade,
		BasicSalary:  salarymatrix.BasicSalary,
		PayCut:       salarymatrix.PayCut,
		Allowance:    salarymatrix.Allowance,
		HeadOfFamily: salarymatrix.HeadOfFamily,
		UpdatedAt:    salarymatrix.UpdatedAt,
	}
}

func (ad *AddSalaryMatrix) NewSalaryMatrix(salarymatrix *AddSalaryMatrix) *SalaryMatrix {
	grade, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", salarymatrix.Grade), 64)
	basicsalary, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", salarymatrix.BasicSalary), 64)
	paycut, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", salarymatrix.PayCut), 64)
	allowance, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", salarymatrix.Allowance), 64)
	headoffamily, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", salarymatrix.HeadOfFamily), 64)

	return &SalaryMatrix{
		Grade:        int(grade),
		BasicSalary:  basicsalary,
		PayCut:       paycut,
		Allowance:    allowance,
		HeadOfFamily: headoffamily,
	}
}

func (sr *AddSalaryMatrix) ValidateFiled() {
	if sr.Grade == nil ||
		sr.BasicSalary == nil ||
		sr.PayCut == nil ||
		sr.Allowance == nil ||
		sr.HeadOfFamily == nil {
		panic(exception.NewBadRequestError("SALARY_MATRIX.NOT_CONTAIN_NEEDED_PROPERTY"))
	}

	if reflect.TypeOf(sr.Grade).Kind() != reflect.Float64 ||
		reflect.TypeOf(sr.BasicSalary).Kind() != reflect.Float64 ||
		reflect.TypeOf(sr.PayCut).Kind() != reflect.Float64 ||
		reflect.TypeOf(sr.Allowance).Kind() != reflect.Float64 ||
		reflect.TypeOf(sr.HeadOfFamily).Kind() != reflect.Float64 {
		panic(exception.NewBadRequestError("SALARY_MATRIX.NOT_MEET_DATA_TYPE_SPECIFICATION"))
	}
}
