package kuji

import (
	"math/rand"
)

type kuji int

const (
	Dikichi   kuji = 1
	Kichi     kuji = 2
	Chuukichi kuji = 3
	Syoukichi kuji = 4
	Suekichi  kuji = 5
	Kyou      kuji = 6
	Daikyou   kuji = 7
)

func (k kuji) PrintFortune() string {
	switch k {
	case 1:
		return "大吉"
	case 2:
		return "吉"
	case 3:
		return "中吉"
	case 4:
		return "小吉"
	case 5:
		return "末吉"
	case 6:
		return "凶"
	case 7:
		return "大凶"
	default:
		return "Error"
	}
}

func RandomFortune() string {
	fortunes := []kuji{
		Dikichi,
		Kichi,
		Chuukichi,
		Syoukichi,
		Suekichi,
		Kyou,
		Daikyou,
	}
	return fortunes[rand.Intn(len(fortunes))].PrintFortune()
}
