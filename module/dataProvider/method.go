package dataProvider

import (
	"api_weather_station/server"
	"api_weather_station/utils"
	"github.com/gin-gonic/gin"
)

type dataNewResponse struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Time        string  `json:"time"`
}

func (d dataNewResponse) TableName() string {
	return "records"
}

func handleDataNew() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从数据库中获取最新的气象数据
		data := &dataNewResponse{}
		err := server.Instance.MySQL.Order("time desc").Limit(1).Find(data).Error
		if err != nil || data.Time == "" {
			c.JSON(utils.HttpBaseResponse(200, 1, "invalid data", nil))
			return
		}
		c.JSON(utils.HttpBaseResponse(200, 0, "ok", data))
	}
}

func handleDataHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(utils.HttpBaseResponse(200, 0, "ok", nil))
	}
}
