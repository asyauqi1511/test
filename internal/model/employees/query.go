package employees

const (
	// SELECT
	queryGetAllEmployees = "SELECT id, first_name, last_name, email, hire_date FROM employees"
	queryGetEmployeeByID = queryGetAllEmployees + " WHERE id=$1"

	// INSERT
	queryInsertEmployee = "INSERT INTO employees (first_name, last_name, email, hire_date) VALUES ($1, $2, $3, $4) RETURNING id"

	// UPDATE
	queryUpdateEmployee = "UPDATE employees SET first_name = $2, last_name = $3, email = $4, hire_date = $5 WHERE id = $1"

	// DELETE
	queryDeleteEmployee = "DELETE FROM employees WHERE id=$1"
)
