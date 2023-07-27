package serviceinterface

import (
	"context"
	"payroll-exam/domain"
)

type PayrollServiceInterface interface {
	Create(ctx context.Context, payroll *domain.AddPayRoll) *domain.AddedPayRoll
	FindByEmployeeId(ctx context.Context, employeeId string) *domain.EmployeeXPayroll
	FindAll(ctx context.Context) []*domain.AddedPayRoll
	FindById(ctx context.Context, payrollId string) *domain.AddedPayRoll
}
