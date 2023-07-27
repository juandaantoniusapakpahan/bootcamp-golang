package database

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"fmt"
	"time"
)

type Student struct {
	FirstName  string
	MiddleName string
	LastName   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func Insert(student *Student, db *sql.DB) {
	query := "insert into students(firstname, midname, lastname, created_at, updated_at) values(?,?,?,?,?)"
	_, err := db.Exec(query, student.FirstName, student.MiddleName, student.LastName, time.Now(), time.Now())
	if err != nil {
		panic(err)
	}

}

func FindByName(firstname, lastname string, db *sql.DB) *Student {
	query := "SELECT * FROM students where firstname=? AND lastname=?"
	rows, err := db.Query(query, firstname, lastname)
	defer rows.Close()
	if err != nil {
		panic(err)
	}
	student := Student{}
	if rows.Next() {
		err := rows.Scan(&student.FirstName, &student.MiddleName, &student.LastName, &student.CreatedAt, &student.UpdatedAt)
		if err != nil {
			panic(err)
		}
	} else {
		return &student
	}
	return &student
}

func DeleteAll(db *sql.DB) {
	_, err := db.Exec("delete from students where 1=1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Berhasil menghapus seluruh data")
}

func MyDB() *sql.DB {
	dbUsername := "root"
	dbPassword := "@root123"
	dbName := "db_test"
	dbPort := "3306"
	dbHost := "localhost"

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)
	return db
}
