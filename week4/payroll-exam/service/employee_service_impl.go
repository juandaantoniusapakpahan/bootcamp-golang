package service

import (
	"context"
	"database/sql"
	"payroll-exam/domain"
	"payroll-exam/helper"
	"payroll-exam/repository/repositoryinterface"
	"payroll-exam/service/serviceinterface"
)

type EmployeeServiceImplement struct {
	EmployeRepository repositoryinterface.EmployeeRepositoryInterface
	DB                *sql.DB
}

func NewEmployeeServiceImplement(employerepository repositoryinterface.EmployeeRepositoryInterface, db *sql.DB) serviceinterface.EmployeeServiceInterface {
	return &EmployeeServiceImplement{
		EmployeRepository: employerepository,
		DB:                db,
	}
}

func (es *EmployeeServiceImplement) Create(ctx context.Context, employee *domain.AddEmployee) *domain.AddedEmployee {
	employee.ValidateFiled()

	tx, err := es.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollBack(tx)
	ep := employee.NewEmployee()

	result := es.EmployeRepository.Add(ctx, tx, ep)

	return result
}

func (es *EmployeeServiceImplement) FindById(ctx context.Context, employeeId string) *domain.AddedEmployee {
	tx, err := es.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollBack(tx)
	employed := es.EmployeRepository.FindById(ctx, tx, employeeId)

	return employed
}

func (es *EmployeeServiceImplement) FindAll(ctx context.Context) []*domain.AddedEmployee {
	tx, err := es.DB.Begin()
	if err != nil {
		panic(err)
	}

	employees := es.EmployeRepository.FindAll(ctx, tx)

	return employees
}
