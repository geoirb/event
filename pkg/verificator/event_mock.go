package verificator

import (
	"github.com/stretchr/testify/mock"
)

// EventMock ...
type EventMock struct {
	mock.Mock
}

// Type ...
func (m *EventMock) Type(t string) bool {
	args := m.Called(t)

	if v, ok := args.Get(0).(bool); ok {
		return v
	}
	return false
}
