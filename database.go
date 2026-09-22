package main

import (
	"database/sql"
	"fmt"
	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB

func connectDB() {
	var err error
	connectionString := "server=localhost;database=employee_db;trusted_connection=yes;encrypt=disable"
	db, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to the database successfully!")
}