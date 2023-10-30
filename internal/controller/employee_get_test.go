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

func TestController_EmployeeGet(t *testing.T) {
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
			// Test if got empty parameter.
			name: "invalid id empty",
			args: func() *gin.Context {
				ctx := gin.Context{}
				return &ctx
			}(),
			mock:           func() {},
			wantHttpStatus: http.StatusBadRequest,
			wantErr:        fmt.Errorf("strconv.ParseInt: parsing \"\": invalid syntax"),
		},
		{
			// Test if got not numeric parameter.
			name: "invalid id not numeric",
			args: func() *gin.Context {
				ctx := gin.Context{
					Params: gin.Params{gin.Param{
						Key:   "id",
						Value: "ABCD",
					}},
				}

				return &ctx
			}(),
			mock:           func() {},
			wantErr:        fmt.Errorf("strconv.ParseInt: parsing \"ABCD\": invalid syntax"),
			wantHttpStatus: http.StatusBadRequest,
		},
		{
			// Test if got <= 0 parameter, which is not suitable for employee id.
			name: "invalid id <= 0",
			args: func() *gin.Context {
				ctx := gin.Context{
					Params: gin.Params{gin.Param{
						Key:   "id",
						Value: "0",
					}},
				}

				return &ctx
			}(),
			mock:           func() {},
			wantErr:        fmt.Errorf("employee id not valid"),
			wantHttpStatus: http.StatusBadRequest,
		},
		{
			// Test if model return error when get employee data.
			name: "model return error",
			args: func() *gin.Context {
				ctx := gin.Context{
					Params: gin.Params{gin.Param{
						Key:   "id",
						Value: "123",
					}},
				}

				return &ctx
			}(),
			mock: func() {
				mockEmployee.EXPECT().GetEmployeeByID(gomock.Any(), int64(123)).Return(entity.Employee{}, fmt.Errorf("sql: failed to select from table employees"))
			},
			wantErr:        fmt.Errorf("sql: failed to select from table employees"),
			wantHttpStatus: http.StatusInternalServerError,
		},
		{
			// Success.
			name: "success",
			args: func() *gin.Context {
				ctx := gin.Context{
					Params: gin.Params{gin.Param{
						Key:   "id",
						Value: "123",
					}},
				}

				return &ctx
			}(),
			mock: func() {
				mockEmployee.EXPECT().GetEmployeeByID(gomock.Any(), int64(123)).Return(entity.Employee{
					ID:        123,
					FirstName: "John",
					LastName:  "Bob",
					Email:     "jhon@example.com",
					HireDate:  time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
				}, nil)
			},
			wantData: RequestEmployee{
				ID:        123,
				FirstName: "John",
				LastName:  "Bob",
				Email:     "jhon@example.com",
				HireDate:  pkg.NewDateTime(time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			// Test if no error but no data from database.
			name: "success but not any data yet",
			args: func() *gin.Context {
				ctx := gin.Context{
					Params: gin.Params{gin.Param{
						Key:   "id",
						Value: "123",
					}},
				}
				return &ctx
			}(),
			mock: func() {
				mockEmployee.EXPECT().GetEmployeeByID(gomock.Any(), int64(123)).Return(entity.Employee{}, nil)
			},
			wantData: RequestEmployee{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			c := &Controller{
				employee: mockEmployee,
			}

			gotHttpStatus, _, gotErr := c.EmployeeGet(tt.args)

			assert.Equal(t, tt.wantHttpStatus, gotHttpStatus)
			if tt.wantErr != nil {
				assert.NotNil(t, gotErr, tt.name)
				assert.Equal(t, tt.wantErr.Error(), gotErr.Error(), tt.name)
			} else {
				assert.Nil(t, gotErr, tt.name)
			}
		})
	}
}
