package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c *Controller) EmployeeDelete(ctx *gin.Context) (httpStatus int, data any, err error) {
	var (
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

	// Call model to delete employee data.
	err = c.employee.DeleteEmployeeByID(ctx, nil, employeeID)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return
}
