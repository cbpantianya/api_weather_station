package baseServer

import (
	"api_weather_station/server"
)

type baseServer struct {
}

var Instance *baseServer

func init() {
	Instance = &baseServer{}
	server.RegisterModule(Instance)
}

func (b baseServer) GetModuleInfo() server.ModuleInfo {
	return server.ModuleInfo{
		ID:       "cbpantianya.baseServer",
		Instance: Instance,
	}
}

func (b baseServer) Init() {

}

func (b baseServer) Serve(server *server.Server) {
	group := server.HttpEngine.Group("/server")
	group.GET("/ping", handlePing())
	group.POST("/statue", handleStatue())
}
