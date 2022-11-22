package providers

import (
	"context"
	"github.com/samber/do"
	"sync"
)

var (
	Injector = do.New()
	Once     sync.Once
)

func Initialize(ctx context.Context) {
	// mysql
	do.ProvideNamed[*MysqlProvider](Injector, "main_mysql", func(i *do.Injector) (*MysqlProvider, error) {
		return NewMysql(i, ctx)
	})
	// mysql mock
	do.ProvideNamed[*MysqlProvider](Injector, "main_mysql_mock", func(i *do.Injector) (*MysqlProvider, error) {
		return NewMysqlMock(i, ctx)
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
