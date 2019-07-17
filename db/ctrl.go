package db

import "net/http"

// HandlerGetAllNews ...
func HandlerGetAllNews(w http.ResponseWriter, r *http.Request) {
	err := getAllNews(w, r)
	if err != nil {
		return
	}
}
