package utils

import (
	"flag"
	"github.com/spf13/viper"
)

func IsTest() bool {
	if flag.Lookup("test.v") != nil {
		return true
	}
	if viper.GetString("MODE") == "test" {
		return true
	}
	return false
}
