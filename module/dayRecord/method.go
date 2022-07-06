package dayRecord

import (
	"api_weather_station/server"
	"api_weather_station/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type RecordRequest struct {
	Temperature float64 `json:"temperature" binding:"required"`
	Humidity    float64 `json:"humidity" binding:"required"`
}

type Record struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Time        int64   `json:"time"`
	SN          string  `json:"sn"`
}

func handleUploadFromStation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request RecordRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(utils.HttpBaseResponse(400, 1, "invalid request", nil))
			return
		} else {
			sn := c.GetHeader("sn")
			// 是的，你没有看错，就只有一个设备
			if sn != "9d0dd9f2e3bc" {
				c.JSON(utils.HttpBaseResponse(400, 1, "invalid request", nil))
				return
			}

			record := Record{
				Temperature: request.Temperature,
				Humidity:    request.Humidity,
				Time:        time.Now().UnixMilli(),
				SN:          sn,
			}

			if err := server.Instance.MySQL.Create(&record).Error; err != nil {
				c.JSON(utils.HttpBaseResponse(400, 1, "invalid request", nil))
				return
			}

			c.JSON(utils.HttpBaseResponse(200, 0, "ok", nil))
			return
		}
	}
}
