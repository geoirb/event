package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/geoirb/event/pkg/model"
)

type verificator interface {
	EventType(t string) bool
}

type eventStorage interface {
	Create(ctx context.Context, e model.Event) error
	Update(ctx context.Context, e model.Event) error
}

type eventBuilder interface {
	NewStartedEvent(t string) model.Event
	NewFinishedEvent(t string) model.Event
}

type service struct {
	verificator verificator
	builder     eventBuilder
	storage     eventStorage

	logger log.Logger
}

// NewEvent service.
func NewEvent(
	verificator verificator,
	builder eventBuilder,
	storage eventStorage,

	logger log.Logger,
) Event {
	return &service{
		verificator: verificator,
		builder:     builder,
		storage:     storage,

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

	e := s.builder.NewStartedEvent(r.Type)
	if err = s.storage.Create(ctx, e); err != nil {
		level.Error(logger).Log("msg", "event create in storage", "err", err)
	}
	return
}

// Finish event.
func (s *service) Finish(ctx context.Context, r FinishEvent) (err error) {
	logger := log.WithPrefix(s.logger, "method", "Finish")

	e := s.builder.NewFinishedEvent(r.Type)
	if err = s.storage.Create(ctx, e); err != nil {
		level.Error(logger).Log("msg", "event create in storage", "err", err)
	}
	return
}
