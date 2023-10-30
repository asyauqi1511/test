package controller

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEmployee := NewMockEmployeeResource(ctrl)

	t.Run("Testing init", func(t *testing.T) {
		New(mockEmployee)
	})
}
