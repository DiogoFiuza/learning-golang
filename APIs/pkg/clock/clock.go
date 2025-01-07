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

// For tests
type FakeClock struct{}

func (c FakeClock) Now() time.Time {
	return time.Date(2023, 6, 30, 20, 0, 0, 0, time.Local)
}
