package app

import (
	"database/sql"
	"golang-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	connStr := "postgres://postgres:16979390749@localhost/mst_category?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
