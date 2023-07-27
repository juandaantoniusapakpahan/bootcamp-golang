package test

import (
	"payroll-exam/domain"
	"payroll-exam/exception"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayrollNotContainNeededProperty(t *testing.T) {
	payrollData := []domain.AddPayRoll{
		{
			Priode: "July-2022",
			Hadir:  20.0,
			Absen:  0.0,
		},
		{
			Priode: "July-2022",
			Hadir:  20.0,
			Absen:  0.0,
		},
		{
			EmployeeId: "123123",
			Priode:     "July-2022",
			Hadir:      20.0,
		},
		{
			EmployeeId: "123123",
			Priode:     "July-2022",
			Hadir:      "20",
		},
	}

	for _, v := range payrollData {
		func() {
			defer func() {
				err := recover()
				assert.NotNil(t, err)
				ext, ok := err.(exception.BadRequestError)
				assert.Equal(t, true, ok)
				assert.Equal(t, "PAYROLL.NOT_CONTAIN_NEEDED_PROPERTY", ext.Error)
			}()
			v.ValidateFiled()
		}()

	}
}

func TestPayrollNotMeetDataTypeSpecification(t *testing.T) {
	payrollData := []domain.AddPayRoll{
		{
			EmployeeId: 123123,
			Priode:     "July-2022",
			Hadir:      20.0,
			Absen:      0.0,
		},
		{
			EmployeeId: "123123",
			Priode:     2022,
			Hadir:      20.0,
			Absen:      0.0,
		},
		{
			EmployeeId: "123123",
			Priode:     "July-2022",
			Hadir:      "20",
			Absen:      0.0,
		},
		{
			EmployeeId: "123123",
			Priode:     "July-2022",
			Hadir:      20.0,
			Absen:      " 0.0",
		},
	}

	for _, v := range payrollData {
		func() {
			defer func() {
				err := recover()
				assert.NotNil(t, err)
				ext, ok := err.(exception.BadRequestError)
				assert.Equal(t, true, ok)
				assert.Equal(t, "PAYROLL.NOT_MEET_DATA_TYPE_SPECIFICATION", ext.Error)
			}()
			v.ValidateFiled()
		}()

	}
}

func TestPayrollValidPayload(t *testing.T) {
	payrollData := []domain.AddPayRoll{
		{
			EmployeeId: "123123",
			Priode:     "July-2022",
			Hadir:      20.0,
			Absen:      0.0,
		},
		{
			EmployeeId: "afsasdfas234",
			Priode:     "July-2022",
			Hadir:      15.0,
			Absen:      5.0,
		},
		{
			EmployeeId: "2394jdnsdfg",
			Priode:     "Maret-2022",
			Hadir:      19.0,
			Absen:      1.0,
		},
	}
	for _, v := range payrollData {
		func() {
			defer func() {
				err := recover()
				assert.Nil(t, err)
			}()
			v.ValidateFiled()
		}()
	}
}
