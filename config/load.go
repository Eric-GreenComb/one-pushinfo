package config

import (
	"strings"

	"github.com/spf13/viper"

	"github.com/Eric-GreenComb/one-pushinfo/bean"
)

// Server Server Config
var Server bean.ServerConfig

// Nsq nsq配置
var Nsq bean.NsqConfig

func init() {
	readConfig()
	initConfig()
}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
}

func initConfig() {
	Server.Port = strings.Split(viper.GetString("server.port"), ",")
	Server.Mode = viper.GetString("server.mode")

	Nsq.Host = viper.GetString("nsq.host")
	Nsq.Topic = viper.GetString("nsq.topic")
	Nsq.TopicNum = viper.GetInt("nsq.topic_num")

}
