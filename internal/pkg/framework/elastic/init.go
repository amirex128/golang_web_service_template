package elastic

import (
	"context"
	"errors"
	"fmt"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/assert"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/initializer"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/safe"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/xlog"
	"sync"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

var (
	client              = make(map[string]*elastic.Client, 0)
	clientLock          = &sync.RWMutex{}
	once                = &sync.Once{}
	elasticConnExpected = make([]elasticExpected, 0)

	all  map[string][]initializer.Simple
	lock sync.RWMutex
)

type elasticExpected struct {
	containerName string
	host          string
	port          int
}

func Initialize(ctx context.Context) {
	once.Do(func() {
		safe.Try(func() error {
			for i := range elasticConnExpected {
				var err error
				connection := fmt.Sprintf("http://%s:%d", elasticConnExpected[i].host, elasticConnExpected[i].port)
				clientLock.Lock()
				client[elasticConnExpected[i].containerName], err = elastic.NewClient(
					elastic.SetURL(connection),
					elastic.SetSniff(false),
					elastic.SetHealthcheck(false),
				)
				if err != nil {
					xlog.GetWithError(ctx, errors.New("connect to elastic failed to ")).Error(err)
					return err
				}
				_, _, err = client[elasticConnExpected[i].containerName].Ping(connection).Do(ctx)
				clientLock.Unlock()
				if err != nil {
					xlog.GetWithError(ctx, errors.New("ping to elastic failed to ")).Error(err)
					return err
				}
				logrus.Infof("successfully connected to elastic : %s", connection)
			}
			return nil
		}, 30*time.Second)
	})
}

func MustGetElasticClient(cnt string) *elastic.Client {
	clientLock.RLock()
	defer clientLock.RUnlock()
	val, ok := client[cnt]
	assert.True(ok)
	assert.NotNil(val)
	return val
}

func RegisterElastic(cnt, host string, port int) error {
	elasticConnExpected = append(elasticConnExpected, elasticExpected{
		containerName: cnt,
		host:          host,
		port:          port,
	})
	return nil
}

// Register a new object to inform it after elastic is loaded
func Register(cnt string, m ...initializer.Simple) {
	lock.Lock()
	all[cnt] = make([]initializer.Simple, 0)
	all[cnt] = append(all[cnt], m...)
	lock.Unlock()
}
