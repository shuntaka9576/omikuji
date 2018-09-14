package omikuji

import (
	"time"
	"github.com/shuntaka9576/omikuji/kuji"
)

type Clock interface {
	Now() time.Time
}

type Omikuji struct {
	Clock Clock
	Kuji []kuji.Kuji
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

func (o *Omikuji) Run() {
	now := o.now()
	d := now.Format("20060102")

	switch d[4:] {

	}
}
