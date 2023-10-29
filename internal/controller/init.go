package controller

import (
	"context"
	"github.com/asyauqi1511/test/internal/entity"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -package=controller -destination=mock_controller_test.go -source=init.go
type EmployeeResource interface {
	GetAllEmployees(ctx context.Context) ([]entity.Employee, error)
	GetEmployeeByID(ctx context.Context, id int64) (entity.Employee, error)
	InsertEmployee(ctx context.Context, tx *sqlx.Tx, data entity.Employee) (int64, error)
	UpdateEmployee(ctx context.Context, tx *sqlx.Tx, data entity.Employee) error
	DeleteEmployeeByID(ctx context.Context, tx *sqlx.Tx, id int64) error
}

type Controller struct {
	employee EmployeeResource
}

func New(emp EmployeeResource) Controller {
	return Controller{
		employee: emp,
	}
}
