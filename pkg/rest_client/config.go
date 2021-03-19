package rest_client

import (
	"log"
	"os"
)

type Config struct {
	url string
}

func NewConfig() *Config {
	var c Config
	var exists bool
	c.url, exists = os.LookupEnv("REST_URL")
	if !exists {
		log.Fatalf("Variable REST_URL is unknown")
	}
	return &c
}
