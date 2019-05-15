package worcker

import (
	"log"

	"github.com/DiegoSantosWS/gocrawler/news"
	"github.com/jasonlvhit/gocron"
)

// Start inicia processo para coletar as informações
func Start() {
	//Execura primeira vez
	news.StartWSITEBRASIL()

	log.Println("start scheduler")
	
	//Criar um crontab que é executado a cada minuto
	gocron.Every(1).Minute().Do(news.StartWSITEBRASIL)
	<-gocron.Start()
}
