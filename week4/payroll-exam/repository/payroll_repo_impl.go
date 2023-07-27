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

type PayrollRepositoryImpolement struct{}

func NewPayrollRepositoryImpolement() repositoryinterface.PayrollRepositoryInterface {
	return &PayrollRepositoryImpolement{}
}

func (pr *PayrollRepositoryImpolement) Add(ctx context.Context, Tx *sql.Tx, payroll *domain.PayRoll) *domain.AddedPayRoll {

	payroll.IdPayRool = "payroll-" + uuid.New().String()
	payroll.CreatedAt = time.Now().Format(time.DateTime)
	payroll.UpdatedAt = time.Now().Format(time.DateTime)
	query := `INSERT INTO payrolls(
		payroll_id, 
		priode, 
		jumlah_hadir, 
		jumlah_absen, 
		basic_salary, 
		pay_cut,
		 additional_salary, 
		 head_of_family,
		 total, 
		 employee_id,
		 created_at, 
		 updated_at)
		 VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`
	_, err := Tx.ExecContext(ctx, query,
		payroll.IdPayRool,
		payroll.Priode,
		payroll.JumlahHadir,
		payroll.JumlahAbsen,
		payroll.BasicSalary,
		payroll.PayCut,
		payroll.AdditionalSalary,
		payroll.HeadOfFamily,
		payroll.Total,
		payroll.EmployeeId,
		payroll.CreatedAt,
		payroll.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err.Error())

		panic(err.Error())
	}
	return domain.NewAddedPayroll(payroll)
}

func (pr *PayrollRepositoryImpolement) FindByEmployeeId(ctx context.Context, Tx *sql.Tx, employeeId string) []domain.PayRollXEmployee {
	query := `SELECT 
	payroll_id, 
	priode, 
	jumlah_hadir, 
	jumlah_absen,
	basic_salary, 
	pay_cut,
	additional_salary,
	head_of_family,
	total,
	created_at,
	updated_at
	from payrolls
	where employee_id = ?`

	rows, err := Tx.QueryContext(ctx, query, employeeId)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	payrolls := []domain.PayRollXEmployee{}

	for rows.Next() {
		payroll := domain.PayRollXEmployee{}
		err := rows.Scan(
			&payroll.IdPayRool,
			&payroll.Priode,
			&payroll.JumlahHadir,
			&payroll.JumlahAbsen,
			&payroll.BasicSalary,
			&payroll.PayCut,
			&payroll.AdditionalSalary,
			&payroll.HeadOfFamily,
			&payroll.Total,
			&payroll.CreatedAt,
			&payroll.UpdatedAt,
		)
		if err != nil {
			panic(err)
		}
		payrolls = append(payrolls, payroll)
	}

	return payrolls
}

func (pr *PayrollRepositoryImpolement) IsPriode(ctx context.Context, Tx *sql.Tx, employeeId, priode string) {
	query := `SELECT * from payrolls where employee_id = ? AND priode = ?`
	rows, err := Tx.QueryContext(ctx, query, employeeId, priode)
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}
	if rows.Next() {
		panic(exception.NewBadRequestError("Employee " + priode + " payroll list has been registered"))
	}
}

func (pr *PayrollRepositoryImpolement) FindAll(ctx context.Context, Tx *sql.Tx) []*domain.AddedPayRoll {
	query := `
	select 
	payroll_id, 
	priode, 
	jumlah_hadir,
	jumlah_absen,
	basic_salary,
	pay_cut,
	additional_salary,
	head_of_family,
	total,
	employee_id,
	created_at,
	updated_at from payrolls`

	rows, err := Tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	payrolls := []*domain.AddedPayRoll{}

	for rows.Next() {
		py := domain.AddedPayRoll{}
		rows.Scan(&py.IdPayRool, &py.Priode, &py.JumlahHadir,
			&py.JumlahAbsen,
			&py.BasicSalary, &py.PayCut, &py.AdditionalSalary,
			&py.HeadOfFamily, &py.Total, &py.EmployeeId,
			&py.CreatedAt, &py.UpdatedAt)

		payrolls = append(payrolls, &py)
	}
	return payrolls
}

func (pr *PayrollRepositoryImpolement) FindById(ctx context.Context, Tx *sql.Tx, payrollId string) *domain.AddedPayRoll {
	query := `SELECT 
	payroll_id, 
	priode, 
	jumlah_hadir, 
	jumlah_absen,
	basic_salary, 
	pay_cut,
	additional_salary,
	head_of_family,
	total,
	created_at,
	updated_at
	from payrolls
	where payroll_id = ?`

	rows, err := Tx.QueryContext(ctx, query, payrollId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	py := domain.AddedPayRoll{}
	if rows.Next() {
		er := rows.Scan(&py.IdPayRool, &py.Priode, &py.JumlahHadir, &py.JumlahAbsen,
			&py.BasicSalary, &py.PayCut, &py.AdditionalSalary, &py.HeadOfFamily, &py.Total,
			&py.CreatedAt, &py.UpdatedAt,
		)
		if er != nil {
			fmt.Println(er.Error())
			panic(er.Error())
		}

	} else {
		panic(exception.NewNotFoundError("payroll not found"))
	}

	return &py
}
