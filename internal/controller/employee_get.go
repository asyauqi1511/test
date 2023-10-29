package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c *Controller) EmployeeGet(ctx *gin.Context) (httpStatus int, data any, err error) {
	var (
		resp       RequestEmployee
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

	// Call model to get employee data.
	entityData, err := c.employee.GetEmployeeByID(ctx, employeeID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	data = resp.FromEntity(entityData)
	return
}
