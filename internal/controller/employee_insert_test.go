package controller

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/asyauqi1511/test/internal/entity"
	"github.com/asyauqi1511/test/internal/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestController_EmployeeInsert(t *testing.T) {
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
			// Test if request body is not valid json format.
			name: "body not valid",
			args: func() *gin.Context {
				ctx := gin.Context{
					Request: httptest.NewRequest(http.MethodPost, "http://example.com/employees/123", bufio.NewReader(bytes.NewReader([]byte("<html>")))),
				}
				return &ctx
			}(),
			mock:           func() {},
			wantErr:        fmt.Errorf("invalid character '<' looking for beginning of value"),
			wantHttpStatus: http.StatusBadRequest,
		},
		{
			// Test if request body empty on required field.
			name: "request body empty on required field",
			args: func() *gin.Context {
				ctx := gin.Context{
					Request: httptest.NewRequest(http.MethodPost, "http://example.com/employees/123", bufio.NewReader(bytes.NewReader([]byte(`
						{
							"first_name": "John",
							"email": "john@example.com",
							"hire_date": "2022-12-31"
						}
					`)))),
				}
				return &ctx
			}(),
			mock:           func() {},
			wantErr:        fmt.Errorf("Key: 'RequestEmployee.LastName' Error:Field validation for 'LastName' failed on the 'required' tag"),
			wantHttpStatus: http.StatusBadRequest,
		},
		{
			// Test if email format is not valid.
			name: "email not valid",
			args: func() *gin.Context {
				ctx := gin.Context{
					Request: httptest.NewRequest(http.MethodPost, "http://example.com/employees/123", bufio.NewReader(bytes.NewReader([]byte(`
						{
							"first_name": "John",
							"last_name": "Bob",
							"email": "john.com",
							"hire_date": "2022-12-31"
						}
					`)))),
				}
				return &ctx
			}(),
			mock:           func() {},
			wantErr:        fmt.Errorf("Key: 'RequestEmployee.Email' Error:Field validation for 'Email' failed on the 'email' tag"),
			wantHttpStatus: http.StatusBadRequest,
		},
		{
			// Test if hire date format is not valid.
			name: "hire date format not valid",
			args: func() *gin.Context {
				ctx := gin.Context{
					Request: httptest.NewRequest(http.MethodPost, "http://example.com/employees/123", bufio.NewReader(bytes.NewReader([]byte(`
						{
							"first_name": "John",
							"last_name": "Bob",
							"email": "john.com",
							"hire_date": "2022-12-31 00:00:00"
						}
					`)))),
				}
				return &ctx
			}(),
			mock:           func() {},
			wantErr:        fmt.Errorf("parsing time \"\\\"2022-12-31 00:00:00\\\"\" as \"\\\"2006-01-02\\\"\": cannot parse \" 00:00:00\\\"\" as \"\\\"\""),
			wantHttpStatus: http.StatusBadRequest,
		},
		{
			// Test if model failed to insert employee data.
			name: "failed to insert employee data",
			args: func() *gin.Context {
				ctx := gin.Context{
					Request: httptest.NewRequest(http.MethodPost, "http://example.com/employees/123", bufio.NewReader(bytes.NewReader([]byte(`
						{
							"first_name": "John",
							"last_name": "Bob",
							"email": "john@example.com",
							"hire_date": "2022-12-31"
						}
					`)))),
				}
				return &ctx
			}(),
			mock: func() {
				mockEmployee.EXPECT().InsertEmployee(gomock.Any(), nil, entity.Employee{
					FirstName: "John",
					LastName:  "Bob",
					Email:     "john@example.com",
					HireDate:  time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
				}).Return(int64(0), fmt.Errorf("sql: erorr insert data to employees"))
			},
			wantErr:        fmt.Errorf("sql: erorr insert data to employees"),
			wantHttpStatus: http.StatusInternalServerError,
		},
		{
			// Success.
			name: "success",
			args: func() *gin.Context {
				ctx := gin.Context{
					Request: httptest.NewRequest(http.MethodPost, "http://example.com/employees/123", bufio.NewReader(bytes.NewReader([]byte(`
						{
							"first_name": "John",
							"last_name": "Bob",
							"email": "john@example.com",
							"hire_date": "2022-12-31"
						}
					`)))),
				}
				return &ctx
			}(),
			mock: func() {
				mockEmployee.EXPECT().InsertEmployee(gomock.Any(), nil, entity.Employee{
					FirstName: "John",
					LastName:  "Bob",
					Email:     "john@example.com",
					HireDate:  time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC),
				}).Return(int64(1), nil)
			},
			wantData: RequestEmployee{
				ID:        1,
				FirstName: "John",
				LastName:  "Bob",
				Email:     "john@example.com",
				HireDate:  pkg.NewDateTime(time.Date(2022, 12, 31, 0, 0, 0, 0, time.UTC)),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			c := &Controller{
				employee: mockEmployee,
			}

			gotHttpStatus, gotData, gotErr := c.EmployeeInsert(tt.args)

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
