package service

import (
	"context"

	"github.com/stretchr/testify/mock"
)

// EventMock ...
type EventMock struct {
	mock.Mock
}

// Start ...
func (m *EventMock) Start(ctx context.Context, r StartEvent) (err error) {
	args := m.Called(r)
	return args.Error(0)
}

// Finish ...
func (m *EventMock) Finish(ctx context.Context, r FinishEvent) (err error) {
	args := m.Called(r)
	return args.Error(0)
}
