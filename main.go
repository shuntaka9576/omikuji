package main

import (
	"net/http"
	"github.com/shuntaka9576/omikuji/handler"
)

func main() {
	http.HandleFunc("/omikuzi", handler.OmikujiHandler)
	http.ListenAndServe(":8080", nil)
}
