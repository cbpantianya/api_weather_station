package server

import (
	"api_weather_station/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMySQL() (*gorm.DB, error) {
	// 初始化数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GlobalConfig.MySQL.User,
		config.GlobalConfig.MySQL.Password,
		config.GlobalConfig.MySQL.Host,
		config.GlobalConfig.MySQL.Port,
		config.GlobalConfig.MySQL.Dbname,
	)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GlobalConfig.Redis.Host, config.GlobalConfig.Redis.Port),
		Password: config.GlobalConfig.Redis.Password,
		DB:       config.GlobalConfig.Redis.DB,
	})
	go func() {
		ctx := context.Background()
		if err := rdb.Ping(ctx).Err(); err != nil {
			log.Panic().Err(err).Msg("redis ping failed")
		}
	}()
	return rdb

}
