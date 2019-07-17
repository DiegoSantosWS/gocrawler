package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/DiegoSantosWS/gocrawler/types"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	Db = Connection()
}

const (
	host       = "172.19.0.2"
	port       = 5432
	dbUserName = "postgres"
	dbPassword = "Postgres2019!"
	dbName     = "crawler"
)

func Connection() *sql.DB {
	var err error
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, dbUserName, dbPassword, dbName)
	dbpg, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Println("[ db connect with database] error: ", err)
		return dbpg
	}
	dbpg.SetMaxIdleConns(100)

	err = dbpg.Ping()
	if err != nil {
		log.Println("[ db connect with database] ping error: ", err)
		return dbpg
	}

	fmt.Println("Successfully connected!")
	return dbpg
}

func selectNews(d *sql.DB) []types.Data {
	news := types.Data{}
	allNews := []types.Data{}
	var err error
	rows, err := d.Query("SELECT pid, title, image, description, link, date FROM urls.news ORDER BY pid DESC")
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&news.ID, &news.Title, &news.Image, &news.Description, &news.Link, &news.Date)
		if err != nil {
			log.Println("[ db connect with database] error: ", err)
			return allNews
		}
		allNews = append(allNews, news)
	}
	return allNews
}

func InsertNews(d *sql.DB, n types.Data) {
	var lastInsertId int
	err := d.QueryRow("INSERT INTO urls.news(title, description, link, image, date) VALUES($1,$2,$3,$4,$5) returning pid;", n.Title, n.Description, n.Link, n.Image, n.Date).Scan(&lastInsertId)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("finish")
}
