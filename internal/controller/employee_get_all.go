package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *Controller) EmployeeGetAll(ctx *gin.Context) (httpStatus int, data any, err error) {
	var (
		resp = []RequestEmployee{}
	)

	// Call model to get employee data.
	entityData, err := c.employee.GetAllEmployees(ctx)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	// Parse data to response.
	for _, v := range entityData {
		respData := RequestEmployee{}
		resp = append(resp, respData.FromEntity(v))
	}
	data = resp

	return
}
