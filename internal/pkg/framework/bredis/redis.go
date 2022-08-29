package bredis

import (
	"backend/internal/pkg/framework/array"
	"backend/internal/pkg/framework/assert"
	"backend/internal/pkg/framework/initializer"
	"backend/internal/pkg/framework/safe"
	"backend/internal/pkg/framework/xlog"
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	cli "github.com/go-redis/redis/v9"
	"github.com/sirupsen/logrus"
)

var (
	once              = &sync.Once{}
	redisClient       = make(map[string]cli.Cmdable, 0)
	redisClientLock   = &sync.RWMutex{}
	redisConnExpected = make([]redisExpected, 0)

	all  = make(map[string][]initializer.Simple)
	lock sync.RWMutex
)

type Kind string

type Manager struct {
	Conn cli.Cmdable
	Ctx  context.Context
}

const (
	normalKind  Kind = "normal"
	clusterKind Kind = "cluster"
)

func (k Kind) IsValid() bool {
	return array.StringInArray(string(k), string(normalKind), string(clusterKind))
}

type redisExpected struct {
	containerName string
	host          string
	kind          Kind
	password      string
	port          int
	database      int
}

func RegisterRedis(cnt, host, password, kind string, port, database int) error {
	if !Kind(kind).IsValid() {
		return fmt.Errorf("redis with kind %s not valid", kind)
	}
	redisConnExpected = append(redisConnExpected, redisExpected{
		containerName: cnt,
		database:      database,
		host:          host,
		port:          port,
		password:      password,
		kind:          Kind(kind),
	})
	return nil
}

func Initialize(ctx context.Context) {
	once.Do(func() {
		for i := range redisConnExpected {
			safe.Try(func() error {
				var err error
				if redisConnExpected[i].kind == clusterKind {
					redisClientLock.Lock()
					redisClient[redisConnExpected[i].containerName] = cli.NewClusterClient(&cli.ClusterOptions{
						Addrs: strings.Split(redisConnExpected[i].host, ","),
					})
					_, err = redisClient[redisConnExpected[i].containerName].Ping(ctx).Result()
					redisClientLock.Unlock()
					if err != nil {
						xlog.GetWithError(ctx, errors.New("ping to redis failed to ")).Error(err)
						return err
					}

				} else {
					redisClientLock.Lock()
					redisClient[redisConnExpected[i].containerName] = cli.NewClient(&cli.Options{
						Addr:     fmt.Sprintf("%s:%d", redisConnExpected[i].host, redisConnExpected[i].port),
						Password: redisConnExpected[i].password,
						DB:       redisConnExpected[i].database,
					})
					_, err := redisClient[redisConnExpected[i].containerName].Ping(ctx).Result()
					redisClientLock.Unlock()
					if err != nil {
						xlog.GetWithError(ctx, errors.New("ping to redis failed to ")).Error(err)
						return err
					}
				}
				lock.RLock()
				dbPostInitial, ok := all[redisConnExpected[i].containerName]
				if ok {
					for j := range dbPostInitial {
						dbPostInitial[j].Initial()
					}
				}
				lock.RUnlock()
				logrus.Infof("successfully connected to redis-%s host=%s", redisConnExpected[i].kind, redisConnExpected[i].host)
				return nil
			}, 30*time.Second)
		}

	})
}

func MustGetRedisConn(cnt string) Manager {
	redisClientLock.RLock()
	defer redisClientLock.RUnlock()
	val, ok := redisClient[cnt]
	assert.True(ok)
	assert.NotNil(val)
	return Manager{
		Conn: val,
		Ctx:  context.Background(),
	}
}

// Register a new object to inform it after redis is loaded
func Register(cnt string, m ...initializer.Simple) {
	lock.Lock()
	all[cnt] = make([]initializer.Simple, 0)
	all[cnt] = append(all[cnt], m...)
	lock.Unlock()
}
