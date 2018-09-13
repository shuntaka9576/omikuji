package omikuji

import (
	"time"
	"fmt"
)

type Clock interface {
	Now() time.Time
}

type Omikuji struct {
	Clock
}

func (o *Omikuji) Now() time.Time {
	return o.Clock.Now()
}

func NewOmikuji() Omikuji{

	return &Omikuji
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time {
	return f()
}

func (o *Omikuji) Run() {
	now := o.Clock.Now()
	fmt.Println(now.Month())
	fmt.Println(now.Date())
}
