package repositoryinterface

import (
	"context"
	"database/sql"
	"payroll-exam/domain"
)

type EmployeeRepositoryInterface interface {
	Add(ctx context.Context, Tx *sql.Tx, employee *domain.Employee) *domain.AddedEmployee
	FindById(ctx context.Context, Tx *sql.Tx, employeeId string) *domain.AddedEmployee
	FindAll(ctx context.Context, Tx *sql.Tx) []*domain.AddedEmployee
}
