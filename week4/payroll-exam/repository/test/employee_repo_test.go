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

var tTableEmploye *tests.EmployeeTableHelper = new(tests.EmployeeTableHelper)

func TestAddEmployee(t *testing.T) {
	// Arrange
	tx, err := myDB.Begin()
	if err != nil {
		panic(err)
	}

	payload := domain.Employee{
		IdEmployee:   "employee-1isdf",
		NameEmployee: "juada",
		Gender:       "L",
		Grade:        1,
		IsMarried:    true,
	}

	newEmployeeRepo := repository.NewEmployeeRepositoryImplement()

	// Action
	result := newEmployeeRepo.Add(context.Background(), tx, &payload)
	tx.Commit()
	// Arrange
	findById := tTableEmploye.FindById(myDB, payload.IdEmployee)

	// assert.Equal(t, true, reflect.Ty(result, findById))
	assert.Equal(t, result.IdEmployee, findById.IdEmployee)
	assert.Equal(t, result.Gender, findById.Gender)
	assert.Equal(t, result.Grade, findById.Grade)
	assert.Equal(t, result.IsMarried, findById.IsMarried)

	tTableEmploye.DeleteAll(myDB)
}

func TestFindEmployeeById(t *testing.T) {
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

	payload := domain.Employee{
		IdEmployee:   "ggp-90us",
		NameEmployee: "test",
		Gender:       "L",
		Grade:        1,
		IsMarried:    true,
	}

	tTableEmploye.Add(myDB, payload)
	newRepo := repository.NewEmployeeRepositoryImplement()

	// Action
	employee := newRepo.FindById(context.Background(), tx, payload.IdEmployee)
	tx.Commit()
	// Assert
	assert.Equal(t, payload.IdEmployee, employee.IdEmployee)
	assert.Equal(t, payload.NameEmployee, employee.NameEmployee)
	assert.Equal(t, payload.Gender, employee.Gender)
	tTableEmploye.DeleteAll(myDB)
}

func TestFindEmployeeByIdNotFound(t *testing.T) {
	//Assert
	defer func() {
		err := recover()
		assert.NotNil(t, err)
		ext, ok := err.(exception.NotFoundError)
		assert.Equal(t, true, ok)
		assert.Equal(t, "employee not found", ext.Error)
	}()

	// Arrange
	tx, err := myDB.Begin()
	if err != nil {
		panic(err)
	}

	newEmployeeRepo := repository.NewEmployeeRepositoryImplement()
	// Action
	newEmployeeRepo.FindById(context.Background(), tx, "unknow")
	tx.Commit()

}

func TestFindAllEmployee(t *testing.T) {
	// Arrange
	tx, err := myDB.Begin()
	if err != nil {
		panic(err)
	}

	sampleData1 := domain.Employee{
		IdEmployee: "id-isdjfsi",
	}
	sampleData2 := domain.Employee{
		IdEmployee: "idasdfa",
	}

	tTableEmploye.Add(myDB, sampleData1)
	tTableEmploye.Add(myDB, sampleData2)

	newEmployeeRepo := repository.NewEmployeeRepositoryImplement()

	// Action
	employees := newEmployeeRepo.FindAll(context.Background(), tx)
	tx.Commit()

	// Assert
	assert.Equal(t, 2, len(employees))
	assert.Equal(t, sampleData1.IdEmployee, employees[0].IdEmployee)
	assert.Equal(t, sampleData2.IdEmployee, employees[1].IdEmployee)

}
