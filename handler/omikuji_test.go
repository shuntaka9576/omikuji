package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shuntaka9576/omikuji/handler"
)

func TestOmikujiHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler.OmikujiHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error("unexpected")
		return
	}

	if res.StatusCode != 200 {
		t.Error("Status code error")
		return
	}
}
