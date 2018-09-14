package handler

import (
	"net/http"
	"github.com/shuntaka9576/omikuji/domain"
	"encoding/json"
)

func OmikujiHandler(w http.ResponseWriter, r *http.Request) {
	res := &domain.Omikuji{"吉"}
	json.NewEncoder(w).Encode(res)
}
