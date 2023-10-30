package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestController_EmployeeDelete(t *testing.T) {
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
			// Test if model return error when delete employee data.
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
				mockEmployee.EXPECT().DeleteEmployeeByID(gomock.Any(), nil, int64(123)).Return(fmt.Errorf("sql: failed to delete from table employees"))
			},
			wantErr:        fmt.Errorf("sql: failed to delete from table employees"),
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
				mockEmployee.EXPECT().DeleteEmployeeByID(gomock.Any(), nil, int64(123)).Return(nil)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			c := &Controller{
				employee: mockEmployee,
			}

			gotHttpStatus, _, gotErr := c.EmployeeDelete(tt.args)

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
