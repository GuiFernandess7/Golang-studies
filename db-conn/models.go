package main

import (
	"database/sql"
	"time"
)

type Employee struct {
	Emp_id	 		int
	First_name	 	sql.NullString
	Last_name		sql.NullString
	Birth_day		time.Time
	Sex				sql.NullString
	Salary			sql.NullFloat64
	Super_id   		sql.NullInt64
	Branch_id		sql.NullInt16
}

