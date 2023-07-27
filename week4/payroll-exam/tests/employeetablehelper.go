package tests

import (
	"database/sql"
	"payroll-exam/domain"
	"time"
)

type EmployeeTableHelper struct{}

func (th *EmployeeTableHelper) Add(DB *sql.DB, employee domain.Employee) {

	if employee.IdEmployee == "" {
		employee.IdEmployee = "employe-11"
	}
	if employee.NameEmployee == "" {
		employee.NameEmployee = "name"
	}
	if employee.Gender == "" {
		employee.Gender = "L"
	}
	if employee.Grade == 0 {
		employee.Grade = 1
	}
	if employee.IsMarried == false {
		employee.IsMarried = true
	}
	query := `INSERT INTO employees (employee_id, name, gender, grade, is_married, created_at, updated_at) VALUES(?,?,?,?,?,?,?)`
	_, err := DB.Exec(query, employee.IdEmployee, employee.NameEmployee, employee.Gender, employee.Grade, employee.IsMarried, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}
}

func (th *EmployeeTableHelper) DeleteAll(DB *sql.DB) {
	query := `DELETE from employees where 1=1`
	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}
}

func (th *EmployeeTableHelper) FindById(DB *sql.DB, employeeId string) domain.AddedEmployee {
	query := "SELECT employee_id, name, gender, grade, is_married, created_at, updated_at from employees where employee_id = ?"

	rows, err := DB.Query(query, employeeId)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	addedEmployee := domain.AddedEmployee{}
	if rows.Next() {
		rows.Scan(&addedEmployee.IdEmployee,
			&addedEmployee.NameEmployee,
			&addedEmployee.Gender,
			&addedEmployee.Grade,
			&addedEmployee.IsMarried,
			&addedEmployee.CreatedAt,
			&addedEmployee.UpdatedAt,
		)
	}
	return addedEmployee
}
