package news

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/DiegoSantosWS/gocrawler/db"
	"github.com/DiegoSantosWS/gocrawler/types"
	"github.com/PuerkitoBio/goquery"
	"github.com/jasonlvhit/gocron"
)

func Start() {
	execute()
	scheduler()
}

func scheduler() {
	log.Println("initialized scheduler")
	gocron.Every(1).Minute().Do(execute)
	<-gocron.Start()
}

// Execute ...
func execute() {
	url := []string{
		"https://www.wsitebrasil.com.br/blog/categoria/comercio-eletronico",
		"https://www.wsitebrasil.com.br/blog/categoria/empreendedorismo",
		"https://www.wsitebrasil.com.br/blog/categoria/marketing-digital",
		"https://www.wsitebrasil.com.br/blog/categoria/tecnologia",
	}
	extract(url)
}

func extract(url []string) {
	for _, u := range url {
		time.Sleep(time.Second * 1)
		processURL(u)
	}
}

func processURL(u string) {
	log.Println(fmt.Sprintf("[wsitebrasil] Executing the url [%s]", u))
	rest, err := http.Get(u)
	if err != nil {
		log.Println(fmt.Sprintf("[news wsitebrasil] error from http. url [%s] err [%v]", u, err))
	}

	defer rest.Body.Close()

	doc, err := goquery.NewDocumentFromReader(rest.Body)
	if err != nil {
		log.Println(fmt.Sprintf("[news wsitebrasil] error from http. url [%s] err [%v]", u, err))
	}
	data := types.Data{}
	doc.Find(".box-news").Each(func(i int, g *goquery.Selection) {
		link, e := g.Find("a").First().Attr("href")
		if !e {
			log.Println(fmt.Sprintf("err to content, link not found [%s]", u))
			return
		}
		l := strings.TrimSpace(link)
		imgs, e := g.Find("img").First().Attr("src")
		if !e {
			log.Println(fmt.Sprintf("err to content, img not found [%s]", u))
			return
		}
		img := strings.TrimSpace(imgs)
		g.Find("h3").Each(func(j int, d *goquery.Selection) {
			data.Date = strings.TrimSpace(d.Find(".data-title").Text())
		})

		title := strings.Split(strings.TrimSpace(g.Find("h3").Text()), " \n                       ")
		data.Title = title[1]
		data.Description = strings.TrimSpace(g.Find("article").Text())
		// log.Println(fmt.Sprintf("data [%s]\n", strings.TrimSpace(data.Description)))
		data.Link = l
		data.Image = img
		saveData(data)
	})
}

func saveData(d types.Data) {
	db.InsertNews(db.Db, d)
}
