package main

import (
	_ "api_weather_station/module/baseServer"
	_ "api_weather_station/module/dayRecord"
	"api_weather_station/server"
)

func main() {
	server.Init()
	server.StartService()
	server.Run()
}
