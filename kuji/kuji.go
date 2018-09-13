package kuji

import "time"

type Kuji interface {
	Date() []time.Time
}

type Daikichi struct {
	Dates []time.Time
}

func (d *Daikichi) Date() []time.Time {}

type Shoukichi struct{}

type Chuukichi struct{}
