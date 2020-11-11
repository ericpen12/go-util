package main

import (
	"github.com/spf13/viper"
	go_utils "go-utils"
	"go.uber.org/zap"
	"log"
)

func initSettings() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./examples/logger/")
	err := viper.ReadInConfig()
	return err
}

func main() {
	err := initSettings()
	if err != nil {
		log.Println("cannot initialize setting, ", err)
		return
	}
	go_utils.Init()
	zap.L().Info("this is an info message")
}
