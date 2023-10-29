package employees

import (
	"database/sql"
	"github.com/asyauqi1511/test/internal/entity"
	"time"
)

type Employee struct {
	ID        int64          `db:"id"`
	FirstName sql.NullString `db:"first_name"`
	LastName  sql.NullString `db:"last_name"`
	Email     sql.NullString `db:"email"`
	HireDate  sql.NullTime   `db:"hire_date"`
}

func (e Employee) ToEntity() entity.Employee {
	var result = entity.Employee{}
	result.ID = e.ID

	firstName, _ := e.FirstName.Value()
	result.FirstName = firstName.(string)

	lastName, _ := e.LastName.Value()
	result.LastName = lastName.(string)

	email, _ := e.Email.Value()
	result.Email = email.(string)

	hireDate, _ := e.HireDate.Value()
	result.HireDate = hireDate.(time.Time)

	return result
}
