package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@root123@tcp(localhost:3306)/db_packet")
	if err != nil {
		panic(err)
	}

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(time.Minute * 3)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(10)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(10)

	return db
}
