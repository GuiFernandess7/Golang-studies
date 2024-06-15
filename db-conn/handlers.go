package main

import (
	"database/sql"
	"log"
	"time"
)

func ListEmployees(db *sql.DB) ([]Employee, error) {
	data := []Employee{}
	rows, err := db.Query("SELECT * FROM employee;")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var branch_id sql.NullInt16
		var first_name, last_name, sex sql.NullString
		var birth_day time.Time
		var salary sql.NullFloat64
		var super_id sql.NullInt64
		var emp_id int

		err := rows.Scan(&emp_id, &first_name, &last_name, &birth_day, &sex, &salary, &super_id, &branch_id)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		employee := Employee{
			Emp_id:     emp_id,
			First_name: first_name,
			Last_name:  last_name,
			Birth_day:  birth_day,
			Sex:        sex,
			Salary:     salary,
			Super_id:   super_id,
			Branch_id:  branch_id,
		}

		data = append(data, employee)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return data, nil
}

