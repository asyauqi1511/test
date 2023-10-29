package controller

import (
	"github.com/asyauqi1511/test/internal/entity"
	"github.com/asyauqi1511/test/internal/pkg"
)

type RequestEmployee struct {
	ID        int64        `json:"id"`
	FirstName string       `json:"first_name" binding:"required"`
	LastName  string       `json:"last_name" binding:"required"`
	Email     string       `json:"email" binding:"required,email"`
	HireDate  pkg.DateTime `json:"hire_date" binding:"required"`
}

func (r RequestEmployee) ToEntity() entity.Employee {
	return entity.Employee{
		ID:        r.ID,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		HireDate:  r.HireDate.Value(),
	}
}

func (r *RequestEmployee) FromEntity(employee entity.Employee) RequestEmployee {
	r.ID = employee.ID
	r.FirstName = employee.FirstName
	r.LastName = employee.LastName
	r.Email = employee.Email
	r.HireDate = pkg.NewDateTime(employee.HireDate)

	return *r
}
