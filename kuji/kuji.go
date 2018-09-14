package kuji

type Kuji int

const (
	Daikichi  = iota
	Shoukichi
	Chuukichi
)

func (c Kuji) GetName() string {
	switch c {
	case 1:
		return "大吉"
	}
	return ""
}

//import "time"
//type Kuji interface {
//	Name() string
//	Explain() string
//	Date() []time.Time
//}
//
//type Daikichi struct {
//	Dates []time.Time
//}
//
//func (d *Daikichi) Name() string {
//
//}
//
//func (d *Daikichi) Date() []time.Time {
//
//}
//
//func (d *Daikichi) Explain() string {
//
//}
//
//
//type Shoukichi struct{}
//
//type Chuukichi struct{}
