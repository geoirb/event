package service

import (
	"errors"
)

// service errors.
var (
	ErrWrongFormatEventType = errors.New("wrong format event type")
	ErrNotFound             = errors.New("not found")
)
