package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	HttpEngine struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
	} `toml:"httpEngine"`
	MySQL struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"passwd"`
		Dbname   string `toml:"db"`
	} `toml:"mysql"`

	Redis struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		DB       int    `toml:"db"`
		Password string `toml:"passwd"`
	} `toml:"redis"`

	Log struct {
		Env  string `toml:"env"`
		Path string `toml:"path"`
	} `toml:"log"`
}

var GlobalConfig Config

func Init() {
	// 读取配置文件
	_, err := toml.DecodeFile("./config/config.toml", &GlobalConfig)
	if err != nil {
		// 此处无日志服务，直接Panic，不会监测配置文件的自动更新
		panic(err)
		return
	}
}
