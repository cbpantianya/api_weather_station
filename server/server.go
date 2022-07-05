package server

import (
	"api_weather_station/config"
	"api_weather_station/utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var Instance *Server

type Server struct {
	HttpEngine *gin.Engine
	MySQL      *gorm.DB
	Redis      *redis.Client
}

// Init 服务初始化
func Init() {
	// 注意初始化顺序
	// 一定要先读取配置文件
	config.Init()
	utils.Init()
	Instance = &Server{}
	Instance.initGin()
	var err error
	Instance.MySQL, err = initMySQL()
	Instance.Redis = initRedis()
	if err != nil {
		log.Logger.Panic().Msg("MySQL init failed")
		return
	}

}

// initGin 初始化gin
func (s *Server) initGin() {
	gin.SetMode(gin.ReleaseMode) // 生产模式
	s.HttpEngine = gin.New()     // 创建一个gin实例
	switch config.GlobalConfig.Log.Env {
	case "dev":
		s.HttpEngine.Use(utils.GinLoggerDevMiddleware()) // 开发模式
		break
	case "prod":
		s.HttpEngine.Use(utils.GinLoggerProdMiddleware()) // 生产模式
		break
	}
	s.HttpEngine.Use(gin.Recovery()) // 异常恢复
	s.HttpEngine.Use(cors.Default()) // 跨域

}

// Run 启动服务
func Run() {
	log.Logger.Info().Msg("HTTP server started")
	log.Logger.Info().Msgf("Listen on %s:%d", config.GlobalConfig.HttpEngine.Host, config.GlobalConfig.HttpEngine.Port)
	if config.GlobalConfig.Log.Env == "prod" {
		fmt.Printf("Listen on %s:%d\n", config.GlobalConfig.HttpEngine.Host, config.GlobalConfig.HttpEngine.Port)
	}
	err := Instance.HttpEngine.Run(fmt.Sprintf("%s:%d", config.GlobalConfig.HttpEngine.Host, config.GlobalConfig.HttpEngine.Port))
	if err != nil {
		log.Logger.Panic().Msg("HTTP server died")
		return
	}
}
