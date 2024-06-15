// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package database

import (
	"context"
	"database/sql"
)

const filterEmployeeBySalary = `-- name: FilterEmployeeBySalary :many
SELECT emp_id, first_name, last_name, birth_day, sex, salary, super_id, branch_id
FROM employee
WHERE salary >= $1 AND salary < $2
`

type FilterEmployeeBySalaryParams struct {
	Min sql.NullInt32
	Max sql.NullInt32
}

func (q *Queries) FilterEmployeeBySalary(ctx context.Context, arg FilterEmployeeBySalaryParams) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, filterEmployeeBySalary, arg.Min, arg.Max)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.EmpID,
			&i.FirstName,
			&i.LastName,
			&i.BirthDay,
			&i.Sex,
			&i.Salary,
			&i.SuperID,
			&i.BranchID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const filterEmployeeBySalaryAndSex = `-- name: FilterEmployeeBySalaryAndSex :many
SELECT emp_id, first_name, last_name, birth_day, sex, salary, super_id, branch_id
FROM employee
WHERE salary >= $1 AND salary < $2
AND ($3::VARCHAR IS NULL OR sex = $3::VARCHAR)
`

type FilterEmployeeBySalaryAndSexParams struct {
	Min sql.NullInt32
	Max sql.NullInt32
	Sex string
}

func (q *Queries) FilterEmployeeBySalaryAndSex(ctx context.Context, arg FilterEmployeeBySalaryAndSexParams) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, filterEmployeeBySalaryAndSex, arg.Min, arg.Max, arg.Sex)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.EmpID,
			&i.FirstName,
			&i.LastName,
			&i.BirthDay,
			&i.Sex,
			&i.Salary,
			&i.SuperID,
			&i.BranchID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const filterEmployeeBySex = `-- name: FilterEmployeeBySex :many
SELECT emp_id, first_name, last_name, birth_day, sex, salary, super_id, branch_id
FROM employee
WHERE sex = $1
`

func (q *Queries) FilterEmployeeBySex(ctx context.Context, sex sql.NullString) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, filterEmployeeBySex, sex)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.EmpID,
			&i.FirstName,
			&i.LastName,
			&i.BirthDay,
			&i.Sex,
			&i.Salary,
			&i.SuperID,
			&i.BranchID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findEmployeeByName = `-- name: FindEmployeeByName :many
SELECT emp_id, first_name, last_name, birth_day, sex, salary, super_id, branch_id FROM employee WHERE first_name = $1
`

func (q *Queries) FindEmployeeByName(ctx context.Context, n sql.NullString) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, findEmployeeByName, n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.EmpID,
			&i.FirstName,
			&i.LastName,
			&i.BirthDay,
			&i.Sex,
			&i.Salary,
			&i.SuperID,
			&i.BranchID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAllEmployees = `-- name: ListAllEmployees :many
SELECT emp_id, first_name, last_name, birth_day, sex, salary, super_id, branch_id FROM employee
`

func (q *Queries) ListAllEmployees(ctx context.Context) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, listAllEmployees)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.EmpID,
			&i.FirstName,
			&i.LastName,
			&i.BirthDay,
			&i.Sex,
			&i.Salary,
			&i.SuperID,
			&i.BranchID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAllEmployeesBy = `-- name: ListAllEmployeesBy :many
SELECT emp_id, first_name, last_name, birth_day, sex, salary, super_id, branch_id FROM employee
ORDER BY emp_id DESC
LIMIT $1
`

func (q *Queries) ListAllEmployeesBy(ctx context.Context, limit int32) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, listAllEmployeesBy, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.EmpID,
			&i.FirstName,
			&i.LastName,
			&i.BirthDay,
			&i.Sex,
			&i.Salary,
			&i.SuperID,
			&i.BranchID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listClients = `-- name: ListClients :many
SELECT client_id, client_name, branch_id FROM client
`

func (q *Queries) ListClients(ctx context.Context) ([]Client, error) {
	rows, err := q.db.QueryContext(ctx, listClients)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Client
	for rows.Next() {
		var i Client
		if err := rows.Scan(&i.ClientID, &i.ClientName, &i.BranchID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}