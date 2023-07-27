package repositorymock

import (
	"context"
	"database/sql"
	"payroll-exam/domain"

	"github.com/stretchr/testify/mock"
)

type SalaryMatrixRepositoryMock struct {
	Mock mock.Mock
}

// Add(ctx context.Context, tx *sql.Tx, salarymatrix domain.SalaryMatrix) *domain.AddedSalaryMatrix

func (spm *SalaryMatrixRepositoryMock) Add(ctx context.Context, tx *sql.Tx, salarymatrix domain.SalaryMatrix) *domain.AddedSalaryMatrix {
	args := spm.Mock.Called(salarymatrix.IdSalary)
	if args[0] == nil {
		return nil
	}

	result := args[0].(domain.AddedSalaryMatrix)
	return &result
}

func (spm *SalaryMatrixRepositoryMock) GetAll(ctx context.Context, tx *sql.Tx) []*domain.AddedSalaryMatrix
func (spm *SalaryMatrixRepositoryMock) FindById(ctx context.Context, tx *sql.Tx, salarymatrixId string) *domain.AddedSalaryMatrix
func (spm *SalaryMatrixRepositoryMock) IsThereGrade(ctx context.Context, tx *sql.Tx, grade int) {
}
func (spm *SalaryMatrixRepositoryMock) Update(ctx context.Context, tx *sql.Tx, salarymatrix *domain.SalaryMatrix) *domain.UpdatedSalaryMatrix
func (spm *SalaryMatrixRepositoryMock) IsSalaryMatrix(ctx context.Context, tx *sql.Tx, salarymatrixId string)
func (smp *SalaryMatrixRepositoryMock) GetByGrade(ctx context.Context, tx *sql.Tx, grade int) *domain.SalaryMatrix
