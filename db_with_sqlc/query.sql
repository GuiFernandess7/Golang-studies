-- name: ListAllEmployees :many
SELECT * FROM employee;

-- name: ListAllEmployeesBy :many
SELECT * FROM employee
ORDER BY emp_id DESC
LIMIT $1;

-- name: ListSalariesASC :many
SELECT *
FROM employee
ORDER BY salary ASC;

-- name: ListSalariesDESC :many
SELECT *
FROM employee
ORDER BY salary DESC;

-- name: FindEmployeeByName :many
SELECT * FROM employee WHERE first_name = @n;

-- name: FilterEmployeeBySalary :many
SELECT *
FROM employee
WHERE salary >= @min AND salary < @max;

-- name: FilterEmployeeBySex :many
SELECT *
FROM employee
WHERE sex = @sex;

-- name: FilterEmployeeBySalaryAndSex :many
SELECT *
FROM employee
WHERE salary >= @min AND salary < @max
AND (@sex::VARCHAR IS NULL OR sex = @sex::VARCHAR);

-- name: ListClients :many
SELECT * FROM client;