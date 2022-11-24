package providers

import (
	"context"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.elastic.co/apm/module/apmlogrus"
	"os"
	"sync"
)

var (
	Injector = do.New()
	Once     sync.Once
)

func Initialize(ctx context.Context) {

	logrus.AddHook(&apmlogrus.Hook{})
	logrus.SetOutput(os.Stdout)
	viper.AutomaticEnv()

	viper.SetDefault("CGO_ENABLED", "1")
	viper.SetDefault("GO111MODULE", "on")
	viper.SetDefault("ELASTIC_APM_ENVIRONMENT", "staging")
	viper.SetDefault("ELASTIC_APM_SERVER_URL", "http://localhost:8200")
	viper.SetDefault("ELASTIC_APM_SERVICE_NAME", "selloora_backend")
	viper.SetDefault("SELLOORA_ELASTIC_HOST", "localhost")
	viper.SetDefault("SELLOORA_ELASTIC_PORT", "9200")
	viper.SetDefault("SELLOORA_MYSQL_MAIN_DB", "selloora")
	viper.SetDefault("SELLOORA_MYSQL_MAIN_HOST", "localhost")
	viper.SetDefault("SELLOORA_MYSQL_MAIN_PASSWORD", "q6766581Amirex")
	viper.SetDefault("SELLOORA_MYSQL_MAIN_PORT", "3306")
	viper.SetDefault("SELLOORA_MYSQL_MAIN_USER", "selloora")
	viper.SetDefault("SELLOORA_REDIS_DB", "1")
	viper.SetDefault("SELLOORA_REDIS_HOST", "localhost")
	viper.SetDefault("SELLOORA_REDIS_PORT", "6379")
	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("SERVER_PORT", "8585")
	viper.SetDefault("SERVER_URL", "http://localhost:8585")

	// mysql
	do.ProvideNamed[*MysqlProvider](Injector, "main_mysql", func(i *do.Injector) (*MysqlProvider, error) {
		return NewMysql(i, ctx)
	})
	// redis
	do.ProvideNamed[*RedisProvider](Injector, "main_redis", func(i *do.Injector) (*RedisProvider, error) {
		return NewRedis(i, ctx)
	})
	// elastic
	do.ProvideNamed[*ElasticProvider](Injector, "main_elastic", func(i *do.Injector) (*ElasticProvider, error) {
		return NewElastic(i, ctx)
	})
	// logrus
	do.Provide[*LogrusProvider](Injector, func(i *do.Injector) (*LogrusProvider, error) {
		return NewLogrus(i, ctx)
	})

}
