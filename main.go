package main

import (
	"net/http"

	"fmt"
	"os"

	"github.com/shuntaka9576/omikuji/handler"
)

func main() {
	http.HandleFunc("/omikuji", handler.OmikujiHandler)
	fmt.Fprintf(os.Stdout, "Web Server is available at http://localhost:8080/omikuji")
	http.ListenAndServe(":8080", nil)
}
