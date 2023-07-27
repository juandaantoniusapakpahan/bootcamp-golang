package test

import (
	"payroll-exam/domain"
	"payroll-exam/exception"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSalaryMatrixNotContainNeededProperty(t *testing.T) {

	salarydata := []domain.AddSalaryMatrix{
		{
			BasicSalary:  100000,
			PayCut:       100000,
			Allowance:    100000,
			HeadOfFamily: 100000,
		},
		{
			Grade:        1,
			PayCut:       100000,
			Allowance:    true,
			HeadOfFamily: 10000,
		},
		{
			Grade:        1,
			BasicSalary:  10000,
			Allowance:    10000,
			HeadOfFamily: 10000,
		},
		{
			Grade:        1,
			BasicSalary:  10000,
			PayCut:       100000,
			HeadOfFamily: 10000,
		},
		{
			Grade:       1,
			BasicSalary: 10000,
			PayCut:      100000,
			Allowance:   10000,
		},
	}

	for _, v := range salarydata {
		func() {
			defer func(t *testing.T) {
				er := recover()
				ext, _ := er.(exception.BadRequestError)
				assert.NotNil(t, ext)
				assert.Equal(t, "SALARY_MATRIX.NOT_CONTAIN_NEEDED_PROPERTY", ext.Error)
			}(t)
			v.ValidateFiled()
		}()
	}
}
func TestSalaryMatrixNotMeetDataType(t *testing.T) {

	salarydata := []domain.AddSalaryMatrix{
		{
			Grade:        "103",
			BasicSalary:  1000.0,
			PayCut:       100000.0,
			Allowance:    10000.0,
			HeadOfFamily: 10000.0,
		},
		{
			Grade:        10,
			BasicSalary:  true,
			PayCut:       100000.0,
			Allowance:    10000.0,
			HeadOfFamily: 10000.0,
		},
		{
			Grade:        10,
			BasicSalary:  10000.0,
			PayCut:       false,
			Allowance:    10000.0,
			HeadOfFamily: 10000.0,
		},
		{
			Grade:        10,
			BasicSalary:  10000.0,
			PayCut:       100000.0,
			Allowance:    "10000",
			HeadOfFamily: 10000.0,
		},
		{
			Grade:        10,
			BasicSalary:  10000.0,
			PayCut:       100000.0,
			Allowance:    10000,
			HeadOfFamily: "10000",
		},
	}

	for _, v := range salarydata {
		func() {
			defer func(t *testing.T) {
				er := recover()
				ext, _ := er.(exception.BadRequestError)
				assert.NotNil(t, ext)

				assert.Equal(t, "SALARY_MATRIX.NOT_MEET_DATA_TYPE_SPECIFICATION", ext.Error)
			}(t)
			v.ValidateFiled()
		}()
	}
}

func TestSalaryMatrixValidPayload(t *testing.T) {

	salary := []domain.AddSalaryMatrix{
		{
			Grade:        1.0,
			BasicSalary:  20000.00,
			PayCut:       100000.0,
			Allowance:    10000.0,
			HeadOfFamily: 10000.0,
		},
		{
			Grade:        3.0,
			BasicSalary:  30000.00,
			PayCut:       400000.0,
			Allowance:    50000.0,
			HeadOfFamily: 20000.0,
		},
	}

	for _, v := range salary {
		func() {
			defer func(t *testing.T) {
				er := recover()
				assert.Nil(t, er)
			}(t)
			v.ValidateFiled()

		}()
	}
}
