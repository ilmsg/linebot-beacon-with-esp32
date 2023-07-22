//go:build ignore

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ilmsg/linebot-beacon-with-esp32/bot"
	"github.com/ilmsg/linebot-beacon-with-esp32/config"
	"github.com/ilmsg/linebot-beacon-with-esp32/handler"
	"github.com/joho/godotenv"
	"golang.ngrok.com/ngrok"
	ngrok_config "golang.ngrok.com/ngrok/config"
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

	ctx := context.Background()
	tun, err := ngrok.Listen(
		ctx,
		ngrok_config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("tunnel created:", tun.URL())

	// http.HandleFunc("/", handler.Handler(linebot))
	http.Serve(tun, handler.Handler(linebot))
}
