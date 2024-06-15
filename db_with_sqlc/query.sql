-- name: ListAllEmployees :many
SELECT * FROM employee;

-- name: ListAllEmployeesBy :many
SELECT * FROM employee
ORDER BY emp_id ASC
LIMIT $1;

-- name: ListClients :many
SELECT * FROM client;

-- name: FindEmployeeByName :many
SELECT * FROM employee WHERE first_name = @n;

