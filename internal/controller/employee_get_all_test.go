package controller

import (
	"fmt"
	"github.com/asyauqi1511/test/internal/entity"
	"github.com/asyauqi1511/test/internal/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestController_EmployeeGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEmployee := NewMockEmployeeResource(ctrl)

	tests := []struct {
		name           string
		args           *gin.Context
		mock           func()
		wantHttpStatus int
		wantData       any
		wantErr        error
	}{
		{
			// Test if model return error when delete employee data.
			name: "model return error",
			args: func() *gin.Context {
				ctx := gin.Context{}
				return &ctx
			}(),
			mock: func() {
				mockEmployee.EXPECT().GetAllEmployees(gomock.Any()).Return([]entity.Employee{}, fmt.Errorf("sql: failed to select from table employees"))
			},
			wantErr:        fmt.Errorf("sql: failed to select from table employees"),
			wantHttpStatus: http.StatusInternalServerError,
		},
		{
			// Success.
			name: "success",
			args: func() *gin.Context {
				ctx := gin.Context{}
				return &ctx
			}(),
			mock: func() {
				mockEmployee.EXPECT().GetAllEmployees(gomock.Any()).Return([]entity.Employee{
					{
						ID:        1,
						FirstName: "John",
						LastName:  "Bob",
						Email:     "jhon@example.com",
						HireDate:  time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        2,
						FirstName: "Chris",
						LastName:  "Evan",
						Email:     "chris@example.com",
						HireDate:  time.Date(2021, 1, 1, 1, 2, 3, 4, time.UTC),
					},
				}, nil)
			},
			wantData: []RequestEmployee{
				{
					ID:        1,
					FirstName: "John",
					LastName:  "Bob",
					Email:     "jhon@example.com",
					HireDate:  pkg.NewDateTime(time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        2,
					FirstName: "Chris",
					LastName:  "Evan",
					Email:     "chris@example.com",
					HireDate:  pkg.NewDateTime(time.Date(2021, 1, 1, 1, 2, 3, 4, time.UTC)),
				},
			},
		},
		{
			// Test if no error but no data from database.
			name: "success but not any data yet",
			args: func() *gin.Context {
				ctx := gin.Context{}
				return &ctx
			}(),
			mock: func() {
				mockEmployee.EXPECT().GetAllEmployees(gomock.Any()).Return([]entity.Employee{}, nil)
			},
			wantData: []RequestEmployee{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			c := &Controller{
				employee: mockEmployee,
			}

			gotHttpStatus, gotData, gotErr := c.EmployeeGetAll(tt.args)

			assert.Equal(t, tt.wantHttpStatus, gotHttpStatus)
			assert.Equal(t, tt.wantData, gotData)
			if tt.wantErr != nil {
				assert.NotNil(t, gotErr, tt.name)
				assert.Equal(t, tt.wantErr.Error(), gotErr.Error(), tt.name)
			} else {
				assert.Nil(t, gotErr, tt.name)
			}
		})
	}
}
