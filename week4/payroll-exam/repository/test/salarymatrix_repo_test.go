package test

import (
	"context"
	"payroll-exam/domain"
	"payroll-exam/exception"
	"payroll-exam/repository"
	"payroll-exam/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// db := db.ConnectDBTest()
	m.Run()
	// defer testtable.DELETEALL(db)
}

var testtable *tests.SalaryMatrixTestTable = new(tests.SalaryMatrixTestTable)

func TestAddSalaryMatrix(t *testing.T) {
	tx, err := myDB.Begin()
	if err != nil {
		panic(err)
	}

	payload := domain.SalaryMatrix{
		Grade:        1.0,
		PayCut:       100000.00,
		BasicSalary:  5000000.00,
		Allowance:    200000.00,
		HeadOfFamily: 1000000.00,
	}

	salaryrepo := repository.NewSalaryMatrixImplement()
	result := salaryrepo.Add(context.Background(), tx, payload)
	tx.Commit()

	checkData := testtable.FindById(myDB, result.IdSalary)
	assert.Equal(t, result.Grade, checkData.Grade)
	assert.Equal(t, result.Allowance, checkData.Allowance)
	assert.Equal(t, result.PayCut, checkData.PayCut)
	assert.Equal(t, result.BasicSalary, checkData.BasicSalary)
	assert.Equal(t, result.HeadOfFamily, checkData.HeadOfFamily)
	testtable.DELETEALL(myDB)

}

func TestGetAllSalaryMatrix(t *testing.T) {
	tx, err := myDB.Begin()
	if err != nil {
		panic(err)
	}

	sampleData := []domain.SalaryMatrix{
		{
			IdSalary:     "as9df",
			Grade:        1,
			BasicSalary:  1000000,
			PayCut:       200000,
			Allowance:    2030000,
			HeadOfFamily: 1000000,
		},
		{
			IdSalary:     "asfas",
			Grade:        2,
			BasicSalary:  1000000,
			PayCut:       200000,
			Allowance:    2030000,
			HeadOfFamily: 1000000,
		},
	}

	testtable.Add(myDB, sampleData[0])
	testtable.Add(myDB, sampleData[1])

	repoSalaryMatrix := repository.NewSalaryMatrixImplement()
	result := repoSalaryMatrix.GetAll(context.Background(), tx)
	tx.Commit()

	assert.Equal(t, 2, len(result))
	assert.Equal(t, sampleData[0].IdSalary, result[0].IdSalary)
	assert.Equal(t, sampleData[1].IdSalary, result[1].IdSalary)
	testtable.DELETEALL(myDB)
}

func TestFindSalaryMatrixById(t *testing.T) {
	// Arrange
	tx, err := myDB.Begin()
	if err != nil {
		panic(err)
	}

	newSalaryRepo := repository.NewSalaryMatrixImplement()

	payload := domain.SalaryMatrix{
		IdSalary:     "salary-iasdf",
		Grade:        1,
		BasicSalary:  50000000.0,
		PayCut:       100000.0,
		Allowance:    200000.0,
		HeadOfFamily: 10000000.0,
	}

	testtable.Add(myDB, payload)
	// Action
	result := newSalaryRepo.FindById(context.Background(), tx, "salary-iasdf")
	tx.Commit()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, payload.IdSalary, result.IdSalary)
	assert.Equal(t, payload.Grade, result.Grade)
	assert.Equal(t, payload.BasicSalary, result.BasicSalary)

	testtable.DELETEALL(myDB)
}

func TestFindSalaryMatrixByInvalidId(t *testing.T) {
	// Arrange
	defer testtable.DELETEALL(myDB)
	defer func(t *testing.T) {
		er := recover()
		ext, ok := er.(exception.NotFoundError)
		assert.Equal(t, true, ok)
		assert.Equal(t, "salarymatrix not found", ext.Error)

	}(t)
	tx, err := myDB.Begin()
	if err != nil {
		panic(err)
	}

	payload := domain.SalaryMatrix{
		IdSalary:     "salary-soidfj",
		Grade:        1.0,
		BasicSalary:  12910923.0,
		PayCut:       12121.0,
		Allowance:    1000000.2,
		HeadOfFamily: 1000000.0,
	}

	newRepo := repository.NewSalaryMatrixImplement()
	testtable.Add(myDB, payload)

	// Action
	_ = newRepo.FindById(context.Background(), tx, "aisdfas")
	tx.Commit()
	// Assert
	t.Errorf("did not panic")
}

func TestIsThereGradeFound(t *testing.T) {
	// Assert
	defer func() {
		err := recover()
		ext, ok := err.(exception.BadRequestError)
		assert.Equal(t, true, ok)
		assert.Equal(t, "grades are registered", ext.Error)
		assert.NotNil(t, err)
		testtable.DELETEALL(myDB)
	}()

	// Arrange
	tx, err := myDB.Begin()
	if err != nil {
		panic(err)
	}

	payroll := domain.SalaryMatrix{
		IdSalary:     "salarymatrix-90sfnasf",
		Grade:        1.0,
		BasicSalary:  50000000.00,
		PayCut:       1000000.0,
		Allowance:    200000,
		HeadOfFamily: 10000000,
	}
	testtable.Add(myDB, payroll)
	newSalaryRepo := repository.NewSalaryMatrixImplement()

	// Acction
	newSalaryRepo.IsThereGrade(context.Background(), tx, payroll.Grade)
	tx.Commit()

}

func TestIsThereGradeNotFound(t *testing.T) {
	// Assert
	defer func() {
		err := recover()
		assert.Nil(t, err)
		testtable.DELETEALL(myDB)
	}()

	// Arrange
	tx, err := myDB.Begin()
	if err != nil {
		panic(err)
	}

	payload := domain.SalaryMatrix{}
	testtable.Add(myDB, payload)

	newSalaryRepo := repository.NewSalaryMatrixImplement()

	//Action
	newSalaryRepo.IsThereGrade(context.Background(), tx, 1)
}

func TestUpdateSalaryMatrix(t *testing.T) {
	// Arrange
	tx, err := myDB.Begin()
	if err != nil {
		panic(err)
	}

	payload := domain.SalaryMatrix{
		IdSalary:     "salary-dsaoifj9",
		Grade:        1,
		BasicSalary:  50000000.0,
		PayCut:       100000.0,
		Allowance:    200000.0,
		HeadOfFamily: 1000000.0,
	}

	payloadUpdate := domain.SalaryMatrix{
		IdSalary:     "salary-dsaoifj9",
		Grade:        1,
		BasicSalary:  30000000.0,
		PayCut:       400000.0,
		Allowance:    500000.0,
		HeadOfFamily: 9000000.0,
	}

	testtable.Add(myDB, payload)
	newRepo := repository.NewSalaryMatrixImplement()

	// Action
	result := newRepo.Update(context.Background(), tx, &payloadUpdate)
	tx.Commit()
	// Assert
	assert.Equal(t, payloadUpdate.IdSalary, result.IdSalary)
	assert.Equal(t, payloadUpdate.BasicSalary, result.BasicSalary)
	assert.Equal(t, payloadUpdate.PayCut, result.PayCut)

	testtable.DELETEALL(myDB)
}

func TestIsSalaryMatrixFound(t *testing.T) {
	// Assert
	defer func() {
		err := recover()
		assert.Nil(t, err)
	}()

	// Arrange
	tx, err := myDB.Begin()
	if err != nil {
		panic(err)
	}

	payload := domain.SalaryMatrix{IdSalary: "salary-owei"}
	testtable.Add(myDB, payload)
	newSalaryRepo := repository.NewSalaryMatrixImplement()

	// Action
	newSalaryRepo.IsSalaryMatrix(context.Background(), tx, payload.IdSalary)
	tx.Commit()
	testtable.DELETEALL(myDB)
}

func TestIsSalaryMatrixNotFound(t *testing.T) {
	// Assert
	defer func() {
		err := recover()
		assert.NotNil(t, err)

		ext, ok := err.(exception.NotFoundError)
		assert.Equal(t, true, ok)
		assert.Equal(t, "salarymatrix not found", ext.Error)
	}()

	// Arrange
	tx, err := myDB.Begin()
	if err != nil {
		panic(err)
	}

	payload := domain.SalaryMatrix{IdSalary: "salary-owei"}
	testtable.Add(myDB, payload)
	newSalaryRepo := repository.NewSalaryMatrixImplement()

	// Action
	newSalaryRepo.IsSalaryMatrix(context.Background(), tx, "sadfasidf")
	tx.Commit()
	testtable.DELETEALL(myDB)
}
