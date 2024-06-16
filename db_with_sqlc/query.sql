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

-- name: EmployeesCount :one
SELECT COUNT(emp_id)
FROM employee;

-- name: GetEmployeeByBranch :many
SELECT employee.emp_id, employee.first_name, employee.last_name, branch.branch_name
FROM employee
JOIN branch
ON branch.branch_id = employee.branch_id
WHERE branch.branch_name = @b_name::VARCHAR;

-- name: ListManagers :many
SELECT employee.emp_id, employee.first_name, branch.branch_name
FROM employee
JOIN branch
ON employee.emp_id = branch.mgr_id;

-- name: ListClients :many
SELECT * FROM client;


