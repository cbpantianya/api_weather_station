package utils

import (
	"api_weather_station/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Init() {
	switch config.GlobalConfig.Log.Env {
	case "dev":
		// 初始化开发日志
		initDevLog()
		break
	case "prod":
		// 初始化生产日志
		initProdLog()
		break
	}
	log.Logger.Info().Msg("Logger started")
}

// WARN: 请勿将该日志配置放到生产环境，因为该日志具有性能问题
func initDevLog() {
	// 初始化开发日志
	LogCmdOutputStyle := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	LogCmdOutputStyle.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%-30s|", i)
	}
	LogCmdOutputStyle.FormatLevel = func(i interface{}) string {
		// You should Upper the string first
		UpperLevel := strings.ToUpper(fmt.Sprintf("%s", i))
		// Give the color to the level
		switch UpperLevel {
		case "DEBUG":
			return fmt.Sprintf("|\033[1;34m%s\033[0m|", UpperLevel)
		case "INFO":
			return fmt.Sprintf("|\033[1;32m%s\033[0m |", UpperLevel)
		case "WARN":
			return fmt.Sprintf("|\033[1;33m%s\033[0m |", UpperLevel)
		case "ERROR":
			return fmt.Sprintf("|\033[1;31m%s\033[0m|", UpperLevel)
		}
		return fmt.Sprintf("|\033[1;30;47m%s\033[0m|", UpperLevel)
	}
	LogCmdOutputStyle.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("\033[1;35m%s:\033[0m", i)
	}

	log.Logger = log.With().Timestamp().Logger().Output(LogCmdOutputStyle)
	log.Warn().Msg("In DEV mode!!!!!!!!")
}

func initProdLog() {
	// 初始化生产日志
	dirPath := config.GlobalConfig.Log.Path
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			panic(err)
			return
		}
	}
	logPath := filepath.Join(dirPath, time.Now().Format("2006-1-2")+".log")
	log.Logger = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).Output(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})
}

func GinLoggerDevMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 计算处理时间
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		// 请求IP
		ip := c.ClientIP()

		// 请求方法
		method := c.Request.Method

		// 请求返回状态码
		statusCode := c.Writer.Status()

		// 请求路径
		requestPath := c.Request.URL.Path

		statusCodeString := ""
		if statusCode == 200 {
			statusCodeString = fmt.Sprintf("\033[1;32m%d\033[0m", statusCode)
		} else {
			statusCodeString = fmt.Sprintf("\033[1;31m%d\033[0m", statusCode)
		}

		log.Logger.Info().Msg(fmt.Sprintf("%s | %3s |%10v|%15s| %s ", method, statusCodeString, latencyTime, ip, requestPath))
	}
}

func GinLoggerProdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 计算处理时间
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		// 请求IP
		ip := c.ClientIP()

		// 请求方法
		method := c.Request.Method

		// 请求返回状态码
		statusCode := c.Writer.Status()

		// 请求路径
		requestPath := c.Request.URL.Path

		log.Logger.Info().Str("method", method).Int("statusCode", statusCode).Str("ip", ip).Str("requestPath", requestPath).Dur("latencyTime", latencyTime).Msg("")
	}
}
