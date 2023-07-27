package repositoryinterface

import (
	"context"
	"database/sql"
	"payroll-exam/domain"
)

type PayrollRepositoryInterface interface {
	Add(ctx context.Context, Tx *sql.Tx, payroll *domain.PayRoll) *domain.AddedPayRoll
	IsPriode(ctx context.Context, Tx *sql.Tx, employeeId string, priode string)
	FindByEmployeeId(ctx context.Context, Tx *sql.Tx, employeeId string) []domain.PayRollXEmployee
	FindAll(ctx context.Context, Tx *sql.Tx) []*domain.AddedPayRoll
	FindById(ctx context.Context, Tx *sql.Tx, payrollId string) *domain.AddedPayRoll
}
