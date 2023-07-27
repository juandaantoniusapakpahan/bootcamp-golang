package service

import (
	"context"
	"database/sql"
	"payroll-exam/domain"
	"payroll-exam/helper"
	"payroll-exam/repository/repositoryinterface"
	"payroll-exam/service/serviceinterface"
)

type SalaryMatrixServiceImplement struct {
	SMRepository repositoryinterface.SalaryMatrixRepositoryInterface
	DB           *sql.DB
}

func NewSalaryMatrixServiceImplement(
	smRepository repositoryinterface.SalaryMatrixRepositoryInterface,
	db *sql.DB,
) serviceinterface.SalaryMatrixServiceInterface {
	return &SalaryMatrixServiceImplement{SMRepository: smRepository, DB: db}
}

func (ss *SalaryMatrixServiceImplement) Create(ctx context.Context, salarymatrix domain.AddSalaryMatrix) *domain.AddedSalaryMatrix {
	salarymatrix.ValidateFiled()

	tx, err := ss.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollBack(tx)

	sm := salarymatrix.NewSalaryMatrix(&salarymatrix)
	ss.SMRepository.IsThereGrade(ctx, tx, sm.Grade)

	return ss.SMRepository.Add(ctx, tx, *sm)
}

func (ss *SalaryMatrixServiceImplement) FindById(ctx context.Context, salarymatrixId string) *domain.AddedSalaryMatrix {

	tx, err := ss.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollBack(tx)

	salaraymatrix := ss.SMRepository.FindById(ctx, tx, salarymatrixId)

	return salaraymatrix
}

func (ss *SalaryMatrixServiceImplement) FindAll(ctx context.Context) (_ []*domain.AddedSalaryMatrix) {
	tx, err := ss.DB.Begin()
	if err != nil {
		panic(tx)
	}
	defer helper.CommitOrRollBack(tx)

	salarymatrixs := ss.SMRepository.GetAll(ctx, tx)

	return salarymatrixs
}

func (ss *SalaryMatrixServiceImplement) Edit(ctx context.Context, salarymatrix domain.AddSalaryMatrix, salarymatrixId string) *domain.UpdatedSalaryMatrix {
	salarymatrix.ValidateFiled()

	tx, err := ss.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollBack(tx)

	sm := salarymatrix.NewSalaryMatrix(&salarymatrix)
	sm.IdSalary = salarymatrixId

	ss.SMRepository.IsSalaryMatrix(ctx, tx, salarymatrixId)
	salarymatrixUpdated := ss.SMRepository.Update(ctx, tx, sm)

	return salarymatrixUpdated
}
