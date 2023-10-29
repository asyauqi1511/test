package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) EmployeeInsert(ctx *gin.Context) (httpStatus int, data any, err error) {
	var (
		request    RequestEmployee
		employeeID int64
	)

	// Parse request.
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	// Call model to insert employee data.
	employeeID, err = c.employee.InsertEmployee(ctx, nil, request.ToEntity())
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	// Return employee id as response.
	request.ID = employeeID
	data = request
	return
}
