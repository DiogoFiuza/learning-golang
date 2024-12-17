package clock

import "time"

type (
	Clock interface {
		Now() time.Time
	}

	clock struct{}
)

func NewClock() Clock {
	return clock{}
}

func (c clock) Now() time.Time {
	return time.Now()
}
