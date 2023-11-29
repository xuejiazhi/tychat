package util

import (
	"github.com/spf13/viper"
	"log"
)

var GlobalCfg ServiceConfig

type ServiceConfig struct {
	Msggw     MsgGw          `yaml:"msggw"`
	Admin_Api AdminApiStruct `yaml:"admin_api"`
}

type MsgGw struct {
	Type string `yaml:"type"`
	Port int    `yaml:"port"`
}

type AdminApiStruct struct {
	Port int `yaml:"port"`
}

// PhraseConfig 解析配置文件
func PhraseConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config file failed, %v", err)
	}
	if err := viper.Unmarshal(&GlobalCfg); err != nil {
		log.Printf("unmarshal config file failed, %v", err)
	}

	log.Printf("%#v", GlobalCfg)
}
