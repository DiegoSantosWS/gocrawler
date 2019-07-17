package router

import (
	"log"
	"net/http"

	"github.com/DiegoSantosWS/gocrawler/db"
	"github.com/GuilhermeCaruso/bellt"
)

func Server() {
	router := bellt.NewRouter()

	router.HandleGroup("/v1",
		router.SubHandleFunc("/allnews", db.HandlerGetAllNews, "GET"),
	)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("algum erro aconteceu", err)
	}
}
