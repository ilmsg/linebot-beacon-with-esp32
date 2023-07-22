//go:build ignore

package main

import (
	"log"
	"net/http"

	"github.com/ilmsg/linebot-beacon-with-esp32/bot"
	"github.com/ilmsg/linebot-beacon-with-esp32/config"
	"github.com/ilmsg/linebot-beacon-with-esp32/handler"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	conf := config.NewConfig()
	linebot, err := bot.NewBot(conf)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler.Handler(linebot))
	if err := http.ListenAndServe(":"+conf.Port, nil); err != nil {
		log.Fatal(err)
	}
}
