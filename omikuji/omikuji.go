package omikuji

import (
	"time"
)

type Clock interface {
	Now() time.Time
}

type Omikuji struct {
	Clock Clock
}

func (o *Omikuji) now() time.Time {
	if o.Clock == nil {
		return time.Now()
	}
	return o.Clock.Now()
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time {
	return f()
}

func (o *Omikuji) Run() string {
	now := o.now()
	d := now.Format("20060102")

	switch d[4:] {
	case "0101":
		fallthrough
	case "0102":
		fallthrough
	case "0103":
		return "大吉"
	}
	return ""
}
