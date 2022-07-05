package dataProvider

import (
	"api_weather_station/server"
)

type DataProvider struct {
}

var Instance *DataProvider

func init() {
	Instance = &DataProvider{}
	server.RegisterModule(Instance)
}

func (d DataProvider) GetModuleInfo() server.ModuleInfo {
	return server.ModuleInfo{
		ID:       "cbpantianya.dataProvider",
		Instance: Instance,
	}
}

func (d DataProvider) Init() {

}

func (d DataProvider) Serve(server *server.Server) {
	group := server.HttpEngine.Group("/data")
	group.GET("/new", handleDataNew())
	group.POST("/history", handleDataHistory())
}
