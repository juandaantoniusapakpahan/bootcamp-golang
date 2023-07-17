package test

import (
	"payroll-exam/domain/payroll"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayrollBadPayRoll(t *testing.T) {
	payload := []payroll.PayRollRequest{
		{
			Hadir: 20.0,
			Absen: 0.0,
		},
		{
			EmployeeId: "employe-1",
			Absen:      0.0,
		}, {
			EmployeeId: "employe-1",
			Hadir:      20.0,
		},
		{
			EmployeeId: 112,
			Hadir:      20.0,
			Absen:      0.0,
		},
		{
			EmployeeId: "employe-1",
			Hadir:      "20.0",
			Absen:      0.0,
		},
		{
			EmployeeId: "employe-1",
			Hadir:      20.0,
			Absen:      false,
		},
	}
	for _, v := range payload {
		err := v.ValidateFiled()
		assert.NotNil(t, err)
	}
}

func TestPayrollValidPayRoll(t *testing.T) {
	payload := []payroll.PayRollRequest{
		{
			EmployeeId: "employe-23",
			Hadir:      20.0,
			Absen:      0.0,
		},
		{
			EmployeeId: "employe-adskno",
			Hadir:      10.0,
			Absen:      10.0,
		},
		{
			EmployeeId: "employe-02jsd",
			Hadir:      19.0,
			Absen:      1.0,
		},
	}

	for _, v := range payload {
		err := v.ValidateFiled()
		assert.Nil(t, err)
	}
}
