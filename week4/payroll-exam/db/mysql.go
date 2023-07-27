package db

import (
	"database/sql"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var (
// 	DRIVER   = os.Getenv("DB_DRIVE")
// 	USER     = os.Getenv("DB_USER")
// 	PASSWORD = os.Getenv("DB_PASSWORD")
// 	HOST     = os.Getenv("DB_HOST")
// 	PORT     = os.Getenv("DB_PORT")
// 	DBNAME   = os.Getenv("DB_NAME")
// )

func ConnectDB() *sql.DB {

	db, err := gorm.Open(mysql.New(mysql.Config{DSN: "root:@root123@tcp(localhost:3306)/db_payroll?parseTime=True"}))

	//db, err := sql.Open("mysql", "root:@root123@tcp(localhost:3306)/db_payroll")
	if err != nil {
		panic(err)
	}

	// db.SetMaxOpenConns(5)
	// db.SetConnMaxLifetime(10)
	// db.SetConnMaxIdleTime(5)

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}

	return sqlDb
}

func ConnectDBTest() *sql.DB {
	godotenv.Load(".env")
	db, err := gorm.Open(mysql.New(mysql.Config{DSN: "root:@root123@tcp(localhost:3306)/db_payroll_test?parseTime=True"}))

	//db, err := sql.Open("mysql", "root:@root123@tcp(localhost:3306)/db_payroll_test")
	if err != nil {
		panic(err)
	}

	sqlDb, err := db.DB()

	if err != nil {
		panic(err)
	}
	// db.SetMaxOpenConns(5)
	// db.SetConnMaxLifetime(10)
	// db.SetMaxIdleConns(10)
	return sqlDb
}
