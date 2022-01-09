package viper

import (
	"fmt"
	"log"
	model "server/internal/config"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Viper struct {
}

func (v *Viper) InitViper() (config *model.Config, err error) {
	fileName := "./config.yaml"
	viper.SetConfigFile(fileName)
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		// panic(err.Error())
		return
	}
	// 配置映射
	config = v.ParseSetting()
	// global.GtConfig = config
	return
}

func (v *Viper) ParseSetting() (config *model.Config) {
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
	})
	viper.WatchConfig()
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Panic(err.Error())
	}
	log.Println("配置读取成功！")
	return
}
