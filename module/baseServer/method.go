package baseServer

import (
	"api_weather_station/server"
	"api_weather_station/utils"
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type RWTest struct {
	TestID     int // 自增
	UpdateTime time.Time
}

func (rw *RWTest) TableName() string {
	return "rw_test"
}

func handlePing() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(utils.HttpBaseResponse(200, 0, "ok", "pong"))
	}
}

type statueResponse struct {
	ToMysql float64 `json:"to_mysql"`
	ToRedis float64 `json:"to_redis"`
}

func handleStatue() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mysql 读写时间测试
		startTime := time.Now()
		// 读
		rwTest := RWTest{
			TestID:     1,
			UpdateTime: startTime,
		}
		server.Instance.MySQL.Create(&rwTest)
		// 写
		server.Instance.MySQL.Model(&rwTest).Where(&RWTest{TestID: 1}).Delete(&rwTest)
		// 结束时间
		endTime := time.Now()
		// 计算时间差
		latencyMysql := endTime.Sub(startTime)

		// Redis 读写时间测试
		startTime = time.Now()
		// 读
		server.Instance.Redis.Set(context.Background(), "rw_test", 10, time.Second*10)
		// 写
		server.Instance.Redis.Del(context.Background(), "rw_test")
		// 结束时间
		endTime = time.Now()
		// 计算时间差
		latencyRedis := endTime.Sub(startTime)

		// 返回结果
		c.JSON(utils.HttpBaseResponse(200, 0, "ok", &statueResponse{
			ToMysql: latencyMysql.Seconds(),
			ToRedis: latencyRedis.Seconds(),
		}))

		return

	}
}
