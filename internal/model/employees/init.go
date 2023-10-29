package employees

import (
	"github.com/jmoiron/sqlx"
)

type Module struct {
	db        *sqlx.DB
	writeStmt map[string]*sqlx.Stmt
	readStmt  map[string]*sqlx.Stmt
}

var writeQuery = map[string]string{
	writeStmtInsertEmployee: queryInsertEmployee,
	writeStmtUpdateEmployee: queryUpdateEmployee,
	writeStmtDeleteEmployee: queryDeleteEmployee,
}

var readQuery = map[string]string{
	readStmtGetAllEmployees: queryGetAllEmployees,
	readStmtGetEmployeeByID: queryGetEmployeeByID,
}

func New(db *sqlx.DB) (Module, error) {
	var (
		writeStatement = map[string]*sqlx.Stmt{}
		readStatement  = map[string]*sqlx.Stmt{}
	)

	// Prepare read statement.
	for key, query := range readQuery {
		stmt, err := db.Preparex(query)
		if err != nil {
			return Module{}, err
		}

		readStatement[key] = stmt
	}

	// Prepare write statement.
	for key, query := range writeQuery {
		stmt, err := db.Preparex(query)
		if err != nil {
			return Module{}, err
		}

		writeStatement[key] = stmt
	}

	return Module{
		db:        db,
		writeStmt: writeStatement,
		readStmt:  readStatement,
	}, nil
}
