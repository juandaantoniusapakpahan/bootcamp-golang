package test

import (
	"payroll-exam/domain/salarymatrix"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSMatrixBadPayload(t *testing.T) {
	payload := []salarymatrix.SalaryMatrixRequest{
		{
			BasicSalary:  1000000.0,
			PayCut:       100000.0,
			Allowance:    150000,
			HeadOfFamily: 1000000,
		},
		{
			Grade:        1.0,
			PayCut:       100000.0,
			Allowance:    150000,
			HeadOfFamily: 1000000,
		},
		{
			Grade:        1.0,
			BasicSalary:  1000000.0,
			Allowance:    150000,
			HeadOfFamily: 1000000,
		},
		{
			Grade:        "1.0",
			BasicSalary:  1000000.0,
			PayCut:       100000.0,
			Allowance:    150000,
			HeadOfFamily: 1000000,
		},
		{
			Grade:        1.0,
			BasicSalary:  "1000000.0",
			PayCut:       100000.0,
			Allowance:    150000,
			HeadOfFamily: 1000000,
		},
		{
			Grade:        true,
			BasicSalary:  1000000.0,
			PayCut:       100000.0,
			Allowance:    "150000",
			HeadOfFamily: 1000000,
		},
	}
	for _, v := range payload {
		err := v.ValidateFiled()
		assert.NotNil(t, err)
	}
}

func TestSMatrixValidPayload(t *testing.T) {
	payroll := []salarymatrix.SalaryMatrixRequest{
		{
			Grade:        1.0,
			BasicSalary:  1000000.0,
			PayCut:       100000.0,
			Allowance:    150000.0,
			HeadOfFamily: 1000000.0,
		},
		{
			Grade:        1.0,
			BasicSalary:  1000000.0,
			PayCut:       100000.0,
			Allowance:    150000.0,
			HeadOfFamily: 1000000.0,
		},
		{
			Grade:        1.0,
			BasicSalary:  1000000.0,
			PayCut:       100000.0,
			Allowance:    150000.0,
			HeadOfFamily: 1000000.0,
		},
		{
			Grade:        1.0,
			BasicSalary:  1000000.0,
			PayCut:       100000.0,
			Allowance:    150000.0,
			HeadOfFamily: 1000000.0,
		},
	}

	for _, v := range payroll {
		err := v.ValidateFiled()
		assert.Nil(t, err)
	}
}
