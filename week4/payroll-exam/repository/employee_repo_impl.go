package repository

import (
	"context"
	"database/sql"
	"fmt"
	"payroll-exam/domain"
	"payroll-exam/exception"
	"payroll-exam/repository/repositoryinterface"
	"time"

	"github.com/google/uuid"
)

type EmployeeRepositoryImplement struct {
}

func NewEmployeeRepositoryImplement() repositoryinterface.EmployeeRepositoryInterface {
	return &EmployeeRepositoryImplement{}
}

func (er *EmployeeRepositoryImplement) Add(ctx context.Context, Tx *sql.Tx, employee *domain.Employee) *domain.AddedEmployee {
	employee.IdEmployee = "employee-" + uuid.New().String()
	employee.CreatedAt = time.Now().Format(time.DateTime)
	employee.UpdatedAt = time.Now().Format(time.DateTime)

	query := `INSERT INTO employees (employee_id, name, gender, grade, is_married, created_at, updated_at) VALUES(?,?,?,?,?,?,?)`
	_, err := Tx.ExecContext(ctx, query, employee.IdEmployee,
		employee.NameEmployee, employee.Gender, employee.Grade, employee.IsMarried,
		employee.CreatedAt, employee.UpdatedAt)

	if err != nil {
		fmt.Println(err.Error())

		panic(err.Error())
	}

	return domain.NewAddedEmployee(employee)
}

func (er *EmployeeRepositoryImplement) FindById(ctx context.Context, Tx *sql.Tx, employeeId string) *domain.AddedEmployee {

	query := `SELECT employee_id, name,gender, grade,is_married,created_at,updated_at from employees where employee_id = ?`
	rows, err := Tx.QueryContext(ctx, query, employeeId)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	if !rows.Next() {
		panic(exception.NewNotFoundError("employee not found"))
	}

	employee := domain.AddedEmployee{}

	rows.Scan(&employee.IdEmployee,
		&employee.NameEmployee,
		&employee.Gender,
		&employee.Grade,
		&employee.IsMarried,
		&employee.CreatedAt,
		&employee.UpdatedAt)
	return &employee
}

func (er *EmployeeRepositoryImplement) FindAll(ctx context.Context, Tx *sql.Tx) []*domain.AddedEmployee {
	query := `SELECT employee_id, name, gender, grade, is_married, created_at, updated_at from employees`
	rows, err := Tx.QueryContext(ctx, query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	employees := []*domain.AddedEmployee{}
	for rows.Next() {
		employee := &domain.AddedEmployee{}
		err := rows.Scan(&employee.IdEmployee,
			&employee.NameEmployee,
			&employee.Gender,
			&employee.Grade,
			&employee.IsMarried,
			&employee.CreatedAt,
			&employee.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}
		employees = append(employees, employee)
	}

	return employees
}
