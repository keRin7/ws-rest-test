package main

import (
	"ws-rest-test/wsApp"

	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

func main() {
	AppConfig := wsApp.NewConfig()
	if err := env.Parse(AppConfig); err != nil {
		logrus.Fatal(err)
	}
	App := wsApp.NewWsApp(AppConfig)
	App.Run()

}
