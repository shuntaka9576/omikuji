package main

import (
	"github.com/shuntaka9576/omikuji/omikuji"
	"github.com/shuntaka9576/omikuji/kuji"
	"fmt"
	"os"
)

func main() {
	var nowOmikuji omikuji.Omikuji
	nowOmikuji.Run()
	printKinchi(kuji.Daikichi)
}

func printKinchi(kuji kuji.Kuji) {
	fmt.Fprintf(os.Stdout, (kuji.GetName()))
}
