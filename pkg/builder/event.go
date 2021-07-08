package builder

import (
	"github.com/geoirb/event/pkg/model"
)

type timeFunc func() int64

// Event builder.
type Event struct {
	startedAt  timeFunc
	finishedAt timeFunc
}

// NewEvent return event builder.
func NewEvent(
	startedAt timeFunc,
	finishedAt timeFunc,
) *Event {
	return &Event{
		startedAt:  startedAt,
		finishedAt: finishedAt,
	}
}

func (b *Event) NewStartedEvent(t string) model.Event {
	return model.Event{
		Type:      t,
		State:     model.Unfinished,
		StartedAt: b.startedAt(),
	}
}

func (b *Event) NewFinishedEvent(t string) model.Event {
	return model.Event{
		Type:       t,
		State:      model.Finished,
		FinishedAt: b.finishedAt(),
	}
}
