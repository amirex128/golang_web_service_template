package providers

import (
	"bytes"
	"context"
	"github.com/amirex128/selloora_backend/internal/providers/kv"
	"github.com/amirex128/selloora_backend/internal/providers/safe"
	"github.com/go-redis/redis/v8"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	apmgoredis "go.elastic.co/apm/module/apmgoredisv8"
	"time"
)

type RedisProvider struct {
	Conn *redis.Client
	Ctx  context.Context
}

func (m *RedisProvider) Set(key, value string, ttl time.Duration) error {
	statusCmd := m.Conn.Set(m.Ctx, key, value, ttl)
	if statusCmd.Err() != nil {
		return statusCmd.Err()
	}
	return nil
}

func (m *RedisProvider) Get(key string) (string, error) {
	statusCmd := m.Conn.Get(m.Ctx, key)
	if statusCmd.Err() != nil {
		return "", statusCmd.Err()
	}
	res, err := statusCmd.Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

func (m *RedisProvider) Del(key string) error {
	statusCmd := m.Conn.Del(m.Ctx, key)
	if statusCmd.Err() != nil {
		return statusCmd.Err()
	}
	return nil
}

func (m *RedisProvider) Do(key string, s kv.Serializable, t time.Duration) error {
	target := &bytes.Buffer{}
	err := s.Encode(target)
	if err != nil {
		return err
	}
	return m.Set(key, target.String(), t)
}

func (m *RedisProvider) Hit(key string, s kv.Serializable) error {
	res, err := m.Get(key)
	if err != nil {
		return err
	}
	buf := bytes.NewBufferString(res)
	return s.Decode(buf)
}

func (m *RedisProvider) Exists(key string) bool {
	res := m.Conn.Exists(m.Ctx, key)
	if res.Err() != nil {
		return false
	}
	r, err := res.Result()
	if err != nil {
		return false
	}
	return r >= 1
}

func NewRedis(i *do.Injector, ctx context.Context) (*RedisProvider, error) {
	redisIns := new(RedisProvider)
	safe.Try(func() error {

		redisIns.Conn = redis.NewClient(&redis.Options{
			Addr:     viper.GetString("SELLOORA_REDIS_HOST") + ":" + viper.GetString("SELLOORA_REDIS_PORT"),
			Password: "",                                // no password set
			DB:       viper.GetInt("SELLOORA_REDIS_DB"), // use default DB
		})
		redisIns.Conn.AddHook(apmgoredis.NewHook())

		_, err := redisIns.Conn.Ping(ctx).Result()
		if err != nil {
			logrus.Errorf("failed to register redis service")
			return err
		}
		redisIns.Ctx = ctx

		return nil
	}, 30*time.Second)
	logrus.Infof("redis service is registered")

	return redisIns, nil
}
