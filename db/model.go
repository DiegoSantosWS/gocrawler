package db

import (
	"encoding/json"
	"errors"
	"net/http"
)

func getAllNews(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return errors.New("method not allowed")
	}
	okStatus(w)
	news := selectNews(Db)
	json.NewEncoder(w).Encode(news)
	return errors.New("kdjsklsajdasdalsdjak")
}

func okStatus(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return
}
