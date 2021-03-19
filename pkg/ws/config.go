package ws

import (
	"log"
	"os"
)

type Config struct {
	url     string
	token   string
	channel string
}

func NewConfig(token string) *Config {
	var c Config
	var exists bool

	c.url, exists = os.LookupEnv("URL")
	if !exists {
		log.Fatalf("Variable URL is unknown")
	}
	c.token = token
	/* 	c.token, exists = os.LookupEnv("TOKEN")
	   	if !exists {
	   		log.Fatalf("Variable TOKEN is unknown")
	   	}
	*/
	c.channel, exists = os.LookupEnv("CHANNEL")
	if !exists {
		log.Fatalf("Variable CHANNEL is unknown")
	}
	return &c
}

func (c *Config) SetChannel(channel string) {
	c.channel = channel
}
