package serviceinterface

import (
	"context"
	"payroll-exam/domain"
)

type EmployeeServiceInterface interface {
	Create(ctx context.Context, employee *domain.AddEmployee) *domain.AddedEmployee
	FindById(ctx context.Context, employeeId string) *domain.AddedEmployee
	FindAll(ctx context.Context) []*domain.AddedEmployee
}
