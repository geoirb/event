package verificator

import (
	"regexp"
)

// Event verificator.
type Event struct {
	r *regexp.Regexp
}

// NewEvent ...
func NewEvent(
	layout string,
) (v *Event, err error) {
	r, err := regexp.Compile(layout)
	if err != nil {
		return
	}
	v = &Event{
		r: r,
	}
	return
}

func (e *Event) Type(t string) bool {
	return e.r.Match([]byte(t))
}
