package config

import (
	"fmt"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/assert"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Initialize init config
func Initialize(prefix string) {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Debug("Config file changed:", e.Name)
	})
}

// SetDefault set default and bind
func SetDefault(key string, value interface{}) {
	assert.Nil(viper.BindEnv(key))
	viper.SetDefault(key, value)
}

// GetIntOrDefault
func GetIntOrDefault(key string, def int) int {
	n := viper.GetInt(key)
	if n != 0 {
		return n
	}
	return def
}
