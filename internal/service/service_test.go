package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestService_SpamMask(t *testing.T) {
	input := Service{}
	in := input.SpamMask("http://serdar.com")
	result := ("http://**********")
	assert.Equal(t, in, result, "failed spamMask")
}

type MockProd struct {
	mock.Mock
}

func (m *MockProd) Produce() ([]string, error) {
	args := m.Called()
	return args.Get(0).([]string), args.Error(1)
}

type MockPresent struct {
	mock.Mock
}

func (m *MockPresent) Present(lines []string) error {
	args := m.Called(lines)
	return args.Error(0)
}

func TestService_Run_Success(t *testing.T) {
	mockProduce := new(MockProd)
	mockPresent := new(MockPresent)

	mockProduce.On("Produce").Return([]string{"http://example.com"}, nil)
	mockPresent.On("Present", []string{"http://***********"}).Return(nil)

	service := NewService(mockProduce, mockPresent)

	err := service.Run()

	assert.NoError(t, err)

	mockProduce.AssertExpectations(t)
	mockPresent.AssertExpectations(t)
}
