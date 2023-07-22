package bot

import (
	"github.com/ilmsg/linebot-beacon-with-esp32/config"
	"github.com/line/line-bot-sdk-go/linebot"
)

func NewBot(config *config.Config) (*linebot.Client, error) {
	return linebot.New(
		config.ChannelSecret,
		config.ChannelAccessToken,
	)
}
