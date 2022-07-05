package dayRecord

import (
	"api_weather_station/server"
)

type DayRecord struct {
}

var Instance *DayRecord

func init() {
	Instance = &DayRecord{}
	server.RegisterModule(Instance)
}

func (d DayRecord) GetModuleInfo() server.ModuleInfo {
	return server.ModuleInfo{
		ID:       "cbpantianya.dayRecord",
		Instance: Instance,
	}
}

func (d DayRecord) Init() {

}

func (d DayRecord) Serve(server *server.Server) {
	group := server.HttpEngine.Group("/record")
	group.GET("/upload", handleUploadFromStation())

}
