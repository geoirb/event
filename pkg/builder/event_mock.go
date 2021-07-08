package builder

import (
	"github.com/stretchr/testify/mock"

	"github.com/geoirb/event/pkg/model"
)

// EventMock ...
type EventMock struct {
	mock.Mock
}

// NewStartedEvent ...
func (m *EventMock) NewStartedEvent(t string) model.Event {
	args := m.Called(t)

	if e, ok := args.Get(0).(model.Event); ok {
		return e
	}
	return model.Event{}
}

// NewFinishedEvent ...
func (m *EventMock) NewFinishedEvent(t string) model.Event {
	args := m.Called(t)

	if e, ok := args.Get(0).(model.Event); ok {
		return e
	}
	return model.Event{}
}
