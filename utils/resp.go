package utils

import "github.com/gin-gonic/gin"

func HttpBaseResponse(httpCode int, code int, msg string, data any) (int, any) {
	return httpCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}
