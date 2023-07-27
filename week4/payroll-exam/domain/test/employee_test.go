package test

import (
	"payroll-exam/domain"
	"payroll-exam/exception"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployeeNotMeetDataType(t *testing.T) {
	employeeData := []domain.AddEmployee{
		{
			NameEmployee: "Juanda",
			Grade:        1,
			Gender:       "L",
			IsMarried:    "true",
		},
		{
			NameEmployee: "Juanda",
			Grade:        1,
			Gender:       "L",
			IsMarried:    1,
		},
		{
			NameEmployee: 0,
			Grade:        1,
			Gender:       "L",
			IsMarried:    true,
		},
	}
	for _, v := range employeeData {
		func() {
			defer func(t *testing.T) {
				er := recover()
				ext, _ := er.(exception.BadRequestError)
				assert.NotNil(t, ext)

				assert.Equal(t, "EMPLOYEE.NOT_MEET_DATA_TYPE_SPECIFICATION", ext.Error)
			}(t)
			v.ValidateFiled()
		}()
	}

}

func TestEmployeeNotContainNeededProperty(t *testing.T) {
	employeeData := []domain.AddEmployee{
		{
			Grade:     1,
			Gender:    "P",
			IsMarried: false,
		},
		{
			NameEmployee: "Antoniuse",
			Gender:       "L",
			IsMarried:    true,
		},
		{
			NameEmployee: "Antoniuse",
			Grade:        2,
			IsMarried:    true,
		},
		{
			NameEmployee: "Antoniuse",
			Grade:        2,
			Gender:       "L",
		},
	}

	for _, v := range employeeData {
		func() {
			defer func(t *testing.T) {
				er := recover()
				ext, _ := er.(exception.BadRequestError)
				assert.NotNil(t, ext)
				assert.Equal(t, "EMPLOYEE.NOT_CONTAIN_NEEDED_PROPERTY", ext.Error)
			}(t)
			v.ValidateFiled()
		}()
	}
}

func TestEmployeeValidPayload(t *testing.T) {
	payload := []domain.AddEmployee{
		{
			NameEmployee: "Juanda",
			Gender:       "L",
			Grade:        1.0,
			IsMarried:    true,
		},
		{
			NameEmployee: "Cici",
			Gender:       "P",
			Grade:        2.0,
			IsMarried:    false,
		},
	}

	for _, v := range payload {
		func() {
			defer func() {
				err := recover()
				assert.Nil(t, err)
			}()
			v.ValidateFiled()
		}()
	}

}
