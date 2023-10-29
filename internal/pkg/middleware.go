package pkg

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type stdResp struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Wrap standardize response and logging from controller.
func Wrap(contoller func(c *gin.Context) (int, any, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		var resp stdResp

		httpStatus, data, err := contoller(c)
		if httpStatus == 0 {
			if err != nil {
				httpStatus = http.StatusInternalServerError
			} else {
				httpStatus = http.StatusOK
			}
		}

		resp.Data = data

		if httpStatus == http.StatusOK {
			resp.Status = "Success"
		} else {
			resp.Status = "Failed"
		}

		if err != nil {
			resp.Message = err.Error()
			log.Printf("Error: %v\n", err)
		}

		c.JSON(httpStatus, resp)
	}
}
