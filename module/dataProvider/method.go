package dataProvider

import (
	"api_weather_station/server"
	"api_weather_station/utils"
	"github.com/gin-gonic/gin"
)

type dataNewResponse struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Time        int64   `json:"time"`
}

func (d dataNewResponse) TableName() string {
	return "records"
}

func handleDataNew() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从数据库中获取最新的气象数据
		data := &dataNewResponse{}
		err := server.Instance.MySQL.Order("time desc").Limit(1).Find(data).Error
		if err != nil || data.Time == 0 {
			c.JSON(utils.HttpBaseResponse(200, 1, "invalid data", nil))
			return
		}
		c.JSON(utils.HttpBaseResponse(200, 0, "ok", data))
	}
}

type dataHistoryBody struct {
	StartTime *int64 `json:"start_time" binding:"required"`
	EndTime   *int64 `json:"end_time" binding:"required"`
}

func handleDataHistory() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := &dataHistoryBody{}
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(utils.HttpBaseResponse(400, 1, "invalid request", nil))
			return
		}
		// 从数据库中获取历史数据
		var records []*dataNewResponse
		err := server.Instance.MySQL.Where("time between ? and ?", requestBody.StartTime, requestBody.EndTime).Find(&records).Error
		if err != nil {
			c.JSON(utils.HttpBaseResponse(200, 1, "invalid data", nil))
			return
		}

		c.JSON(utils.HttpBaseResponse(200, 0, "ok", records))
	}
}
