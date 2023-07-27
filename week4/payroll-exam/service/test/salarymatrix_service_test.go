package test

import (
	"context"
	"payroll-exam/domain"
	"payroll-exam/repository/repositorymock"
	"payroll-exam/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockSalarymatrixRepository = &repositorymock.SalaryMatrixRepositoryMock{Mock: mock.Mock{}}
var testservice = service.SalaryMatrixServiceImplement{SMRepository: mockSalarymatrixRepository}

func TestAddService(t *testing.T) {
	// Arrange
	payroll := domain.AddSalaryMatrix{
		// IdSalary:     "1",
		Grade:        1,
		BasicSalary:  5000000,
		PayCut:       500000,
		Allowance:    5000000,
		HeadOfFamily: 1000000,
		// CreatedAt:    "2023-07-25",
		// UpdatedAt:    "2023-07-24",
	}
	//categoryRepository.Mock.On("FindById", "2").Return(dataTest)
	salarymatrix := payroll.NewSalaryMatrix(&payroll)
	salarymatrix.IdSalary = "1"
	salarymatrix.CreatedAt = "2023-07-25"
	salarymatrix.UpdatedAt = "2023-07-24"

	mockSalarymatrixRepository.Mock.On("FindByGrade", mock.Anything, mock.Anything, 1)

	mockSalarymatrixRepository.Mock.On("Add", mock.Anything, 1).Return(payroll)

	result := testservice.Create(context.Background(), payroll)

	assert.NotNil(t, result)
}
