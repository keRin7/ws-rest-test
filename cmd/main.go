package main

import "ws-rest-test/wsApp"

func main() {
	AppConfig := wsApp.NewConfig()
	App := wsApp.NewWsApp(AppConfig)
	App.Run()

}
