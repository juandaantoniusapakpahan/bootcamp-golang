package tests

import (
	"database/sql"
	"payroll-exam/domain"
	"time"
)

type SalaryMatrixTestTable struct {
}

func (st *SalaryMatrixTestTable) Add(db *sql.DB, sm domain.SalaryMatrix) {
	if sm.IdSalary == "" {
		sm.IdSalary = "salary-daidsf"
	}
	if sm.Grade == 0 {
		sm.Grade = 1
	}
	if sm.BasicSalary == 0 {
		sm.Grade = 10000000
	}
	if sm.PayCut == 0 {
		sm.PayCut = 100000
	}
	if sm.Allowance == 0 {
		sm.Allowance = 200000
	}
	if sm.HeadOfFamily == 0 {
		sm.HeadOfFamily = 1000000
	}
	query := `insert into salarymatrixs(
		salary_matrix_id, 
		grade, basic_salary, 
		pay_cut,allowance,
		head_of_family, 
		created_at, 
		updated_at)
		VALUES (?,?,?,?,?,?,?,?)`
	_, err := db.Exec(
		query,
		sm.IdSalary,
		sm.Grade,
		sm.BasicSalary,
		sm.PayCut,
		sm.Allowance,
		sm.HeadOfFamily,
		time.Now().Format(time.DateTime),
		time.Now().Format(time.DateTime),
	)
	if err != nil {
		panic(err)
	}
}

func (st *SalaryMatrixTestTable) DELETEALL(db *sql.DB) {
	query := `DELETE FROM salarymatrixs where 1=1`
	db.Exec(query)
}

func (st *SalaryMatrixTestTable) FindById(db *sql.DB, salarymatrixId string) domain.AddedSalaryMatrix {
	query := "SELECT salary_matrix_id, grade, basic_salary,pay_cut, allowance, head_of_family from salarymatrixs where salary_matrix_id = ?"
	rows, err := db.Query(query, salarymatrixId)
	defer rows.Close()
	if err != nil {

		panic(err)
	}

	sm := domain.AddedSalaryMatrix{}
	if rows.Next() {
		err := rows.Scan(&sm.IdSalary, &sm.Grade, &sm.BasicSalary, &sm.PayCut, &sm.Allowance, &sm.HeadOfFamily)
		if err != nil {
			panic(err)
		}
	} else {
		return sm
	}
	return sm
}
