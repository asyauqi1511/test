package employees

import (
	"context"
	"github.com/asyauqi1511/test/internal/entity"
	"github.com/jmoiron/sqlx"
)

func (m Module) GetAllEmployees(ctx context.Context) ([]entity.Employee, error) {
	var (
		rows   []Employee
		result []entity.Employee
	)

	// Get data from database.
	err := m.readStmt[readStmtGetAllEmployees].SelectContext(ctx, &rows)
	if err != nil {
		return result, err
	}

	// Parse data to entity format.
	for _, row := range rows {
		result = append(result, row.ToEntity())
	}

	return result, nil
}

func (m Module) GetEmployeeByID(ctx context.Context, id int64) (entity.Employee, error) {
	var (
		row Employee
	)

	// Get data from database.
	err := m.readStmt[readStmtGetEmployeeByID].GetContext(ctx, &row, id)
	if err != nil {
		return entity.Employee{}, err
	}

	return row.ToEntity(), nil
}

func (m Module) InsertEmployee(ctx context.Context, tx *sqlx.Tx, data entity.Employee) (int64, error) {
	var (
		employeeID int64
		err        error
	)

	// Create tx if it doesn't exist.
	if tx == nil {
		tx_, errTx := m.db.BeginTxx(ctx, nil)
		if errTx != nil {
			return employeeID, errTx
		}

		tx = tx_
		defer func() {
			if err != nil {
				_ = tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
	}

	// Insert data into database.
	err = m.writeStmt[writeStmtInsertEmployee].QueryRowContext(ctx, data.FirstName, data.LastName, data.Email, data.HireDate).Scan(&employeeID)
	if err != nil {
		return employeeID, err
	}

	return employeeID, nil
}

func (m Module) UpdateEmployee(ctx context.Context, tx *sqlx.Tx, data entity.Employee) error {
	var (
		err error
	)

	// Create tx if it doesn't exist.
	if tx == nil {
		tx_, errTx := m.db.BeginTxx(ctx, nil)
		if errTx != nil {
			return errTx
		}

		tx = tx_
		defer func() {
			if err != nil {
				_ = tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
	}

	// Update data on database.
	_, err = m.writeStmt[writeStmtUpdateEmployee].ExecContext(ctx, data.ID, data.FirstName, data.LastName, data.Email, data.HireDate)
	return err
}

func (m Module) DeleteEmployeeByID(ctx context.Context, tx *sqlx.Tx, id int64) error {
	var (
		err error
	)

	// Create tx if it doesn't exist.
	if tx == nil {
		tx_, errTx := m.db.BeginTxx(ctx, nil)
		if errTx != nil {
			return errTx
		}

		tx = tx_
		defer func() {
			if err != nil {
				_ = tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
	}

	// Update data on database.
	_, err = m.writeStmt[writeStmtDeleteEmployee].ExecContext(ctx, id)
	return err
}
