package employee

import (
	"fmt"
	"strconv"
)

// Employee → IdEmployee, NameEmployee, Gender, Grade, IsMarried

type Employee struct {
	IdEmployee   string
	NameEmployee string
	Gender       string
	Grade        int
	IsMarried    bool
}

type EmployeeInterface interface {
	Add(name string, gender string, grade int, married bool)
	ShowAll()
	FindEmplById(emploId string) Employee
}

type ListEmployee struct {
	employess []Employee
}

func NewEmployee() EmployeeInterface {
	return &ListEmployee{}
}

func (l *ListEmployee) Add(name string, gender string, grade int, married bool) {

	id := "employ-" + strconv.Itoa(len(l.employess))
	newEmployee := Employee{
		IdEmployee:   id,
		NameEmployee: name,
		Gender:       gender,
		Grade:        grade,
		IsMarried:    married,
	}
	l.employess = append(l.employess, newEmployee)
}

func (l *ListEmployee) ShowAll() {
	for _, v := range l.employess {
		fmt.Printf("EmployeId: %s, Name: %s, Gender: %s, Grade: %d, IsMarried: %t", v.IdEmployee, v.NameEmployee, v.Gender, v.Grade, v.IsMarried)
		fmt.Println()
	}
}

func (l *ListEmployee) FindEmplById(emploId string) Employee {
	for _, v := range l.employess {
		if v.IdEmployee == emploId {
			return v
		}
	}
	return Employee{}
}
