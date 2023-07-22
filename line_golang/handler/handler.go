package handler

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

func Handler(bot *linebot.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello from ngork-go")
	}
}
