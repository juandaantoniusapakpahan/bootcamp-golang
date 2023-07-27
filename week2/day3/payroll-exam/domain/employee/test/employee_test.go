package test

import (
	"payroll-exam/domain/employee"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployeeValidPayload(t *testing.T) {

	payload := []employee.EmployeeRequest{
		{
			NameEmployee: "Juanda Antonius Pakpahan",
			Gender:       "L",
			Grade:        1.0,
			IsMarried:    true,
		},
		{
			NameEmployee: "Renatarin",
			Gender:       "P",
			Grade:        1.0,
			IsMarried:    true,
		},
	}

	for _, v := range payload {
		err := v.ValidateFiled()
		assert.Nil(t, err)
	}
}

func TestEmployeeBadPayload(t *testing.T) {
	payload := []employee.EmployeeRequest{
		{
			NameEmployee: 1,
			Gender:       "L",
			Grade:        1.0,
			IsMarried:    true,
		},
		{
			NameEmployee: "Juanda Antonius Pakpahan",
			Gender:       "K",
			Grade:        1.0,
			IsMarried:    true,
		},
		{
			NameEmployee: "Juanda Antonius Pakpahan",
			Gender:       "L",
			Grade:        "i",
			IsMarried:    true,
		},
		{
			NameEmployee: "Juanda Antonius Pakpahan",
			Gender:       "L",
			Grade:        1.0,
			IsMarried:    "true",
		},
		{
			Gender:    "P",
			Grade:     1.0,
			IsMarried: true,
		},
		{
			NameEmployee: "Juanda Antonius Pakpahan",
			Grade:        1.0,
			IsMarried:    true,
		},
		{
			NameEmployee: "Juanda Antonius Pakpahan",
			Gender:       "p",
			IsMarried:    true,
		},
		{
			NameEmployee: "Juanda Antonius Pakpahan",
			Gender:       "L",
			Grade:        1.0,
		},
	}

	for _, v := range payload {
		err := v.ValidateFiled()
		assert.NotNil(t, err)
	}
}
