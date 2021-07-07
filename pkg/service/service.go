package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/geoirb/event/pkg/model"
)

type verificator interface {
	EventType(t string) bool
}

type eventStorage interface {
	Create(ctx context.Context, e model.Event) error
	Update(ctx context.Context, e model.Event)
}

type service struct {
	storage     eventStorage
	verificator verificator

	logger log.Logger
}

// NewEvent service.
func NewEvent(
	storage eventStorage,
	verificator verificator,

	logger log.Logger,
) Event {
	return &service{
		storage:     storage,
		verificator: verificator,

		logger: logger,
	}
}

// Star new event.
func (s *service) Start(ctx context.Context, r StartEvent) (err error) {
	logger := log.WithPrefix(s.logger, "method", "Start")

	if !s.verificator.EventType(r.Type) {
		err = ErrWrongFormatEventType
		level.Error(logger).Log("err", err)
		return
	}

	e := model.Event{
		Type:      r.Type,
		State:     model.Unfinished,
		StartedAt: time.Now().Unix(),
	}

	if err = s.storage.Create(ctx, e); err != nil {
		level.Error(logger).Log("msg", "event create in storage", "err", err)
	}
	return
}

// Finish event.
func (s *service) Finish(ctx context.Context, r FinishEvent) (err error) {
	return
}
