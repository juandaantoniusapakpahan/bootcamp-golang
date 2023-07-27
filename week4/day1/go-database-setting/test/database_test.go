package test

import (
	"database/sql"
	"fmt"
	"go-database-setting/database"
	"testing"
)

var (
	dbUsername = "root"
	dbPassword = "@root123"
	dbName     = "db_test"
	dbPort     = "3306"
	dbHost     = "localhost"
)

func BenchmarkTestMaxOpenCon5(b *testing.B) {

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(5)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			database.Insert(&database.Student{FirstName: "A", MiddleName: "B", LastName: "C"}, db)
		}
	})
}

func BenchmarkTestMaxOpenCon10(b *testing.B) {

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			database.Insert(&database.Student{FirstName: "A", MiddleName: "B", LastName: "C"}, db)
		}
	})

}

func BenchmarkTestMaxOpenCon15(b *testing.B) {

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			database.Insert(&database.Student{
				FirstName:  "A",
				MiddleName: "B",
				LastName:   "C",
			},
				db)
		}
	})

}

func BenchmarkSetConnMaxLifetim5(b *testing.B) {

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(5)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = database.FindByName("A", "C", db)
		}
	})
}
func BenchmarkSetConnMaxLifetim10(b *testing.B) {

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(10)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = database.FindByName("A", "C", db)
		}
	})
}
