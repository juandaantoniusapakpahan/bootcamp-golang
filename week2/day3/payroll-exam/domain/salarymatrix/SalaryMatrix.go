package salarymatrix

import (
	"errors"
	"payroll-exam/db"
	"reflect"

	"github.com/google/uuid"
)

// Struct
// SalaryMatrix → IdSalary, Grade, BasicSalary, PayCut, Allowance, HeadOfFamily

// Task No 4
//Terdapat Fitur menmpilkan daftar matrix salary

type SalaryMatrixRequest struct {
	Grade        interface{} `json:"grade"`
	BasicSalary  interface{} `json:"basic_salary"`
	PayCut       interface{} `json:"pay_cut"`
	Allowance    interface{} `json:"allowance"`
	HeadOfFamily interface{} `json:"head_of_family"`
}

type SalaryMatrix struct {
	IdSalary     string  `json:"salary_matrix_id"`
	Grade        int     `json:"grade"`
	BasicSalary  float64 `json:"basic_salary"`
	PayCut       float64 `json:"pay_cut"`
	Allowance    float64 `json:"allowance"`
	HeadOfFamily float64 `json:"head_of_family"`
}

type SalaryMatrixList struct {
	SalaryMatrixs []SalaryMatrix `json:"salary_matrix"`
}

type SalaryImplement struct {
	DB db.DBInterface
}

type SalaryMatrixInterface interface {
	Add(SalaryMatrix) (SalaryMatrix, error)
	GetAll(channel chan []SalaryMatrix)
	FindByGrade(grade int) SalaryMatrix
	ReadFilePush(newSalarymatrix SalaryMatrix)
}

func NewSalaryMatrix(db db.DBInterface) SalaryMatrixInterface {
	return &SalaryImplement{
		DB: db,
	}
}

func (s *SalaryImplement) ReadFilePush(newSalarymatrix SalaryMatrix) {
	dataSalary := SalaryMatrixList{}
	s.DB.ReadFile("salarymatrix.json", &dataSalary)

	dataSalary.SalaryMatrixs = append(dataSalary.SalaryMatrixs, newSalarymatrix)

	s.DB.MarshalMan("salarymatrix.json", &dataSalary)
}

func (s *SalaryImplement) Add(salarymatrix SalaryMatrix) (SalaryMatrix, error) {
	isSalary := s.FindByGrade(salarymatrix.Grade)
	if isSalary != (SalaryMatrix{}) {
		return SalaryMatrix{}, errors.New("Grade sudah terdaftar")
	}

	id := "matrix-" + uuid.New().String()
	newSalarymatrix := SalaryMatrix{
		IdSalary:     id,
		Grade:        salarymatrix.Grade,
		BasicSalary:  salarymatrix.BasicSalary,
		PayCut:       salarymatrix.PayCut,
		Allowance:    salarymatrix.Allowance,
		HeadOfFamily: salarymatrix.HeadOfFamily,
	}
	s.ReadFilePush(newSalarymatrix)
	return newSalarymatrix, nil
}

func (s *SalaryImplement) GetAll(channel chan []SalaryMatrix) {
	dataSalary := SalaryMatrixList{}
	s.DB.ReadFile("salarymatrix.json", &dataSalary)
	channel <- dataSalary.SalaryMatrixs

}

func (s *SalaryImplement) FindByGrade(grade int) SalaryMatrix {
	listMatrix := SalaryMatrixList{}
	s.DB.ReadFile("salarymatrix.json", &listMatrix)
	for _, v := range listMatrix.SalaryMatrixs {
		if v.Grade == grade {
			return v
		}
	}
	return SalaryMatrix{}
}

func (sr *SalaryMatrixRequest) ValidateFiled() error {
	if sr.Grade == nil ||
		sr.BasicSalary == nil ||
		sr.PayCut == nil ||
		sr.Allowance == nil ||
		sr.HeadOfFamily == nil {
		return errors.New("Mohon lengkapi data")
	}

	if reflect.TypeOf(sr.Grade).Kind() != reflect.Float64 ||
		reflect.TypeOf(sr.BasicSalary).Kind() != reflect.Float64 ||
		reflect.TypeOf(sr.PayCut).Kind() != reflect.Float64 ||
		reflect.TypeOf(sr.Allowance).Kind() != reflect.Float64 ||
		reflect.TypeOf(sr.HeadOfFamily).Kind() != reflect.Float64 {
		return errors.New("Tipe data tidak sesuai")
	}
	return nil
}
