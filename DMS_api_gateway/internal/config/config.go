package config

import (
	"flag"

	"github.com/spf13/viper"
)

func InitConfig() {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	switch res {
	case "":
		viper.AddConfigPath("configs")
		viper.SetConfigName("config")
		if err := viper.ReadInConfig(); err != nil {
			panic("failed to read config: " + err.Error())
		}
	case "test":
		viper.AddConfigPath("configs")
		viper.SetConfigName("config_test")
		if err := viper.ReadInConfig(); err != nil {
			panic("failed to read config: " + err.Error())
		}
	}
}
