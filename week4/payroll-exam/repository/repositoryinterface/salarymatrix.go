package repositoryinterface

import (
	"context"
	"database/sql"
	"payroll-exam/domain"
)

type SalaryMatrixRepositoryInterface interface {
	Add(ctx context.Context, tx *sql.Tx, salarymatrix domain.SalaryMatrix) *domain.AddedSalaryMatrix
	GetAll(ctx context.Context, tx *sql.Tx) []*domain.AddedSalaryMatrix
	FindById(ctx context.Context, tx *sql.Tx, salarymatrixId string) *domain.AddedSalaryMatrix
	IsThereGrade(ctx context.Context, tx *sql.Tx, grade int)
	Update(ctx context.Context, tx *sql.Tx, salarymatrix *domain.SalaryMatrix) *domain.UpdatedSalaryMatrix
	IsSalaryMatrix(ctx context.Context, tx *sql.Tx, salarymatrixId string)
	GetByGrade(ctx context.Context, tx *sql.Tx, grade int) *domain.SalaryMatrix
}
