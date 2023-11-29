package util

import "github.com/spf13/viper"

type ServiceConfig struct {
	Msggw MsgGw `yaml:"msggw"`
}

type MsgGw struct {
	Type string `yaml:"type"`
	Port int    `yaml:"port"`
}

// PhraseConfig 解析配置文件
func PhraseConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("apps")
	viper.SetConfigType("yaml")
}
