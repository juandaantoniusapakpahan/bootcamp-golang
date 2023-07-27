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

type SalaryMatrixRepositoryImplement struct {
}

func NewSalaryMatrixImplement() repositoryinterface.SalaryMatrixRepositoryInterface {
	return &SalaryMatrixRepositoryImplement{}
}

func (sr *SalaryMatrixRepositoryImplement) Add(ctx context.Context, tx *sql.Tx, sm domain.SalaryMatrix) *domain.AddedSalaryMatrix {
	sm.IdSalary = "salarymatrix-" + uuid.New().String()
	sm.CreatedAt = time.Now().Format(time.DateTime)
	sm.UpdatedAt = time.Now().Format(time.DateTime)

	query := `INSERT INTO salarymatrixs (salary_matrix_id, grade, basic_salary, pay_cut, allowance, head_of_family, created_at, updated_at) values (?,?,?,?,?,?,?,?)`

	_, err := tx.ExecContext(ctx, query, sm.IdSalary, sm.Grade, sm.BasicSalary, sm.PayCut, sm.Allowance, sm.HeadOfFamily, sm.CreatedAt, sm.UpdatedAt)
	if err != nil {
		panic(err.Error())
	}
	return domain.NewAddedSalaryMatrix(&sm)

}

func (sr *SalaryMatrixRepositoryImplement) GetAll(ctx context.Context, tx *sql.Tx) (_ []*domain.AddedSalaryMatrix) {
	query := "select salary_matrix_id, grade, basic_salary, pay_cut, allowance, head_of_family, created_at, updated_at from salarymatrixs"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	listsalarymatrix := []*domain.AddedSalaryMatrix{}
	for rows.Next() {
		sm := domain.AddedSalaryMatrix{}
		err := rows.Scan(&sm.IdSalary, &sm.Grade, &sm.BasicSalary, &sm.PayCut, &sm.Allowance, &sm.HeadOfFamily, &sm.CreatedAt, &sm.UpdatedAt)
		if err != nil {
			panic(err)
		}
		listsalarymatrix = append(listsalarymatrix, &sm)
	}

	return listsalarymatrix
}

func (sr *SalaryMatrixRepositoryImplement) FindById(ctx context.Context, tx *sql.Tx, salarymatrixId string) *domain.AddedSalaryMatrix {
	query := `select salary_matrix_id, 
	grade, basic_salary, pay_cut, 
	allowance, head_of_family, created_at,
	 updated_at from salarymatrixs
	where salary_matrix_id = ?`
	rows, err := tx.QueryContext(ctx, query, salarymatrixId)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	sm := domain.AddedSalaryMatrix{}
	if rows.Next() {
		err := rows.Scan(&sm.IdSalary, &sm.Grade, &sm.BasicSalary, &sm.PayCut, &sm.Allowance, &sm.HeadOfFamily, &sm.CreatedAt, &sm.UpdatedAt)
		if err != nil {
			panic(err)
		}
	} else {
		panic(exception.NewNotFoundError("salarymatrix not found"))
	}
	return &sm
}

func (sr *SalaryMatrixRepositoryImplement) IsThereGrade(ctx context.Context, tx *sql.Tx, grade int) {
	query := `select salary_matrix_id, grade, basic_salary, pay_cut, allowance, head_of_family, created_at, updated_at from salarymatrixs
	where grade = ?`
	rows, err := tx.QueryContext(ctx, query, grade)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		panic(exception.NewBadRequestError("grades are registered"))
	}
}

func (sr *SalaryMatrixRepositoryImplement) Update(ctx context.Context, tx *sql.Tx, salarymatrix *domain.SalaryMatrix) (_ *domain.UpdatedSalaryMatrix) {
	query := `update salarymatrixs set grade = ?, 
	basic_salary= ?, pay_cut =?, allowance=?, head_of_family=?, updated_at = ? 
	where salary_matrix_id = ?`
	salarymatrix.UpdatedAt = time.Now().Format(time.DateTime)

	_, err := tx.ExecContext(ctx,
		query,
		salarymatrix.Grade,
		salarymatrix.BasicSalary,
		salarymatrix.PayCut,
		salarymatrix.Allowance,
		salarymatrix.HeadOfFamily,
		salarymatrix.UpdatedAt,
		salarymatrix.IdSalary,
	)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	return domain.NewUpdatedSalaryMatrix(salarymatrix)
}

func (sr *SalaryMatrixRepositoryImplement) IsSalaryMatrix(ctx context.Context, tx *sql.Tx, salaraymatrixId string) {
	query := `select * from salarymatrixs where salary_matrix_id = ?`
	rows, err := tx.QueryContext(ctx, query, salaraymatrixId)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	if !rows.Next() {
		panic(exception.NewNotFoundError("salarymatrix not found"))
	}
}

func (sr *SalaryMatrixRepositoryImplement) GetByGrade(ctx context.Context, tx *sql.Tx, grade int) *domain.SalaryMatrix {
	query := `select salary_matrix_id, grade, basic_salary, pay_cut, allowance, head_of_family, created_at, updated_at from salarymatrixs
	where grade = ?`
	rows, err := tx.QueryContext(ctx, query, grade)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if !rows.Next() {
		panic(exception.NewNotFoundError("salarymatrix not found"))
	}
	salarymatrix := domain.SalaryMatrix{}
	err = rows.Scan(&salarymatrix.IdSalary,
		&salarymatrix.Grade,
		&salarymatrix.BasicSalary,
		&salarymatrix.PayCut,
		&salarymatrix.Allowance,
		&salarymatrix.HeadOfFamily,
		&salarymatrix.CreatedAt,
		&salarymatrix.UpdatedAt,
	)

	return &salarymatrix
}
