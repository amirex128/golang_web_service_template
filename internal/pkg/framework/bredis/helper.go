package bredis

import (
	"bytes"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/kv"
	"time"
)

func (m Manager) Set(key, value string, ttl time.Duration) error {
	statusCmd := m.Conn.Set(m.Ctx, key, value, ttl)
	if statusCmd.Err() != nil {
		return statusCmd.Err()
	}
	return nil
}

func (m Manager) Get(key string) (string, error) {
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

func (m Manager) Del(key string) error {
	statusCmd := m.Conn.Del(m.Ctx, key)
	if statusCmd.Err() != nil {
		return statusCmd.Err()
	}
	return nil
}

// Do save key in redis
func (m Manager) Do(key string, s kv.Serializable, t time.Duration) error {
	target := &bytes.Buffer{}
	err := s.Encode(target)
	if err != nil {
		return err
	}
	return m.Set(key, target.String(), t)
}

// Hit get key from redis
func (m Manager) Hit(key string, s kv.Serializable) error {
	res, err := m.Get(key)
	if err != nil {
		return err
	}
	buf := bytes.NewBufferString(res)
	return s.Decode(buf)
}

// Hit get key from redis
func (m Manager) Exists(key string) bool {
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
