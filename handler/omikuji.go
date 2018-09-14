package handler

import (
	"encoding/json"
	"net/http"

	"github.com/shuntaka9576/omikuji/domain"
	"github.com/shuntaka9576/omikuji/omikuji"
)

func OmikujiHandler(w http.ResponseWriter, r *http.Request) {
	nowOmikuji := omikuji.Omikuji{}

	res := &domain.Omikuji{Result: nowOmikuji.Run()}
	json.NewEncoder(w).Encode(res)
}
