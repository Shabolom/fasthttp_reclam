package main

import (
	adr "advertising_service/cmd/server"
	"github.com/oschwald/geoip2-golang"
	"log"
)

func main() {
	// производим подключение к файлу базы данных
	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := adr.NewServer(db)

	if err := s.Listen(); err != nil {
		log.Panic(err)
	}
}
