package main

import (
	"github.com/DiegoSantosWS/gocrawler/news"
	"github.com/DiegoSantosWS/gocrawler/router"
)

func main() {
	go func() {
		router.Server()
	}()
	news.Start()
}
