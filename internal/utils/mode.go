package utils

import (
	"github.com/spf13/viper"
)

type Mode = string

func GetMode() string {
	mode := viper.GetString("APP_MODE")
	if mode == "" {
		viper.SetDefault("APP_MODE", "debug")
		mode = "debug"
	}
	return mode
}
func IsTest() bool {
	return GetMode() == "test"
}
