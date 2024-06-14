package clock

import "time"

type Clock interface {
	Now() time.Time
}

func NewReal() Clock {
	return realClock{}
}

type realClock struct{}

func (realClock) Now() time.Time { return time.Now() }

func NewFake(now time.Time) FakeClock {
	return FakeClock{now}
}

type FakeClock struct {
	now time.Time
}

func (f FakeClock) Now() time.Time {
	return f.now
}
