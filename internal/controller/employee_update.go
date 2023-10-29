package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c *Controller) EmployeeUpdate(ctx *gin.Context) (httpStatus int, data any, err error) {
	var (
		request    RequestEmployee
		employeeID int64
	)

	// Parse parameter.
	idStr := ctx.Param("id")
	employeeID, err = strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	if employeeID == 0 {
		return http.StatusBadRequest, nil, fmt.Errorf("employee id not valid")
	}

	// Get employee by id to validate data.
	emp, err := c.employee.GetEmployeeByID(ctx, employeeID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	if emp.ID == 0 {
		return http.StatusBadRequest, nil, fmt.Errorf("employee id not found")
	}

	// Parse request.
	err = json.NewDecoder(ctx.Request.Body).Decode(&request)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	// Combine request and existing data.
	if request.FirstName != "" {
		emp.FirstName = request.FirstName
	}
	if request.LastName != "" {
		emp.LastName = request.LastName
	}
	if request.Email != "" {
		emp.Email = request.Email
	}
	if !request.HireDate.IsZero() {
		emp.HireDate = request.HireDate.Value()
		emp.HireDate.IsZero()
	}

	// Call model to update employee data.
	err = c.employee.UpdateEmployee(ctx, nil, emp)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	data = request.FromEntity(emp)
	return
}
