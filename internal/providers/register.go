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
