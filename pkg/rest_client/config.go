package rest_client

//import (
//	"log"
//	"os"
//)

type Config struct {
	Url string `env:"REST_URL"`
}

func NewConfig() *Config {
	return &Config{}
}
