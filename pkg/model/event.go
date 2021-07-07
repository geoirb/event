package model

// State of event type.
type State int

// enum of event states.
const (
	Unfinished State = iota + 1
	Finished
)

// Event ...
type Event struct {
	Type       string
	State      State
	StartedAt  int64
	FinishedAt int64
}
