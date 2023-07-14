package salarymatrix

import (
	"fmt"
	"strconv"
)

// Struct
// SalaryMatrix → IdSalary, Grade, BasicSalary, PayCut, Allowance, HeadOfFamily

// Task No 4
//Terdapat Fitur menmpilkan daftar matrix salary

type SalaryMatrix struct {
	IdSalary     string
	Grade        int
	BasicSalary  float64
	PayCut       float64
	Allowance    float64
	HeadOfFamily float64
}

type SalaryMatrixList struct {
	SalaryMatrixs []SalaryMatrix
}

type SalaryMatrixInterface interface {
	Add(grade int, basicsalary float64, paycut float64, allowance float64, headoffamily float64)
	GetAll()
	FindByGrade(grade int) SalaryMatrix
}

// SalaryMatrix → IdSalary, Grade, BasicSalary, PayCut, Allowance, HeadOfFamily

func NewSalaryMatrix() SalaryMatrixInterface {
	return &SalaryMatrixList{}
}

// type SalaryMatrix struct {
// 	IdSalary     string
// 	Grade        int
// 	BasicSalary  float32
// 	PayCut       float32
// 	Allowance    float32
// 	HeadOfFamily float32
// }

func (s *SalaryMatrixList) Add(grade int, basicsalary float64, paycut float64, allowance float64, headoffamily float64) {
	id := "matrix-" + strconv.Itoa(len(s.SalaryMatrixs))
	newSalarymatrix := SalaryMatrix{
		IdSalary:     id,
		Grade:        grade,
		BasicSalary:  basicsalary,
		PayCut:       paycut,
		Allowance:    allowance,
		HeadOfFamily: headoffamily,
	}
	s.SalaryMatrixs = append(s.SalaryMatrixs, newSalarymatrix)
}

func (s *SalaryMatrixList) GetAll() {
	fmt.Println("==== Salary Matrix ====")
	fmt.Println("MatrixId\t Grade\t BasicSalary\t PayCut\t\t Allowance\t HeadOfFamily")
	for _, v := range s.SalaryMatrixs {
		fmt.Printf("%s\t %d\t %.2f\t %.2f\t %.2f\t %.2f\t \n", v.IdSalary, v.Grade, v.BasicSalary, v.PayCut, v.Allowance, v.HeadOfFamily)
	}
}

func (s *SalaryMatrixList) FindByGrade(grade int) SalaryMatrix {
	for _, v := range s.SalaryMatrixs {
		if v.Grade == grade {
			return v
		}
	}
	return SalaryMatrix{}
}
