package serviceinterface

import (
	"context"
	"payroll-exam/domain"
)

type SalaryMatrixServiceInterface interface {
	Create(ctx context.Context, salarymatrix domain.AddSalaryMatrix) *domain.AddedSalaryMatrix
	FindById(ctx context.Context, salarymatrixId string) *domain.AddedSalaryMatrix
	FindAll(ctx context.Context) []*domain.AddedSalaryMatrix
	Edit(cxt context.Context, salarymatrix domain.AddSalaryMatrix, salarymatrixId string) *domain.UpdatedSalaryMatrix
}
