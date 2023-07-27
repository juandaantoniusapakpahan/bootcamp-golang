package service

import (
	"context"
	"database/sql"
	"payroll-exam/domain"
	"payroll-exam/helper"
	"payroll-exam/repository/repositoryinterface"
	"payroll-exam/service/serviceinterface"
)

type PayrollServiceImplement struct {
	DB                 *sql.DB
	PayrollRepository  repositoryinterface.PayrollRepositoryInterface
	EmployeeRepository repositoryinterface.EmployeeRepositoryInterface
	SM                 repositoryinterface.SalaryMatrixRepositoryInterface
}

func NewPayrollServiceImplement(
	db *sql.DB,
	pr repositoryinterface.PayrollRepositoryInterface,
	er repositoryinterface.EmployeeRepositoryInterface,
	sm repositoryinterface.SalaryMatrixRepositoryInterface,
) serviceinterface.PayrollServiceInterface {
	return &PayrollServiceImplement{
		DB:                 db,
		PayrollRepository:  pr,
		EmployeeRepository: er,
		SM:                 sm,
	}
}

func (ps *PayrollServiceImplement) Create(ctx context.Context, payroll *domain.AddPayRoll) *domain.AddedPayRoll {

	payroll.ValidateFiled()

	tx, err := ps.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollBack(tx)
	py := payroll.NewPayroll()

	ps.PayrollRepository.IsPriode(ctx, tx, py.EmployeeId, py.Priode)
	empl := ps.EmployeeRepository.FindById(ctx, tx, py.EmployeeId)
	sm := ps.SM.GetByGrade(ctx, tx, empl.Grade)

	py.EmployeeId = empl.IdEmployee
	py.BasicSalary = sm.BasicSalary
	py.AdditionalSalary = float64(py.JumlahHadir) * sm.Allowance
	py.PayCut = float64(py.JumlahAbsen) * sm.PayCut
	if empl.Gender == "L" && empl.IsMarried == true {
		py.HeadOfFamily = sm.HeadOfFamily
	}
	py.Total = py.BasicSalary + py.AdditionalSalary - py.PayCut + py.HeadOfFamily

	result := ps.PayrollRepository.Add(ctx, tx, py)

	return result
}

func (ps *PayrollServiceImplement) FindByEmployeeId(ctx context.Context, employeeId string) *domain.EmployeeXPayroll {
	tx, err := ps.DB.Begin()
	defer helper.CommitOrRollBack(tx)
	if err != nil {
		panic(err.Error())
	}

	empl := ps.EmployeeRepository.FindById(ctx, tx, employeeId)

	pys := ps.PayrollRepository.FindByEmployeeId(ctx, tx, employeeId)

	pyl := domain.EmployeeXPayroll{
		IdEmployee:   employeeId,
		NameEmployee: empl.NameEmployee,
		Gender:       empl.Gender,
		Grade:        empl.Grade,
		IsMarried:    empl.IsMarried,
		CreatedAt:    empl.CreatedAt,
		UpdatedAt:    empl.UpdatedAt,
		Payrolls:     pys,
	}

	return &pyl
}

func (ps *PayrollServiceImplement) FindAll(ctx context.Context) []*domain.AddedPayRoll {
	tx, err := ps.DB.Begin()
	defer helper.CommitOrRollBack(tx)
	if err != nil {
		panic(err)
	}
	payrolls := ps.PayrollRepository.FindAll(ctx, tx)
	return payrolls
}

func (ps *PayrollServiceImplement) FindById(ctx context.Context, payrollId string) *domain.AddedPayRoll {
	tx, err := ps.DB.Begin()
	defer helper.CommitOrRollBack(tx)
	if err != nil {
		panic(err.Error())
	}
	payroll := ps.PayrollRepository.FindById(ctx, tx, payrollId)

	return payroll
}
