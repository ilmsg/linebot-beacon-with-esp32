package config

import (
	"os"
)

type Config struct {
	ChannelSecret      string
	ChannelAccessToken string

	Port string
}

func NewConfig() *Config {
	return &Config{
		ChannelSecret:      os.Getenv("CHANNEL_SECRET"),
		ChannelAccessToken: os.Getenv("CHANNEL_ACCESSTOKEN"),
		Port:               os.Getenv("PORT"),
	}
}
