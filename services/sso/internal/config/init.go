package config

import "github.com/spf13/viper"

func Init() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("sso-config")

	return viper.ReadInConfig()
}
