package service

import "github.com/stretchr/testify/mock"

type MockService struct {
	mock.Mock
}

func (m *MockService) Healthz() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockService) FooBar() string {
	args := m.Called()
	return args.String(0)
}
