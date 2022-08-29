package api

import (
	"backend/internal/pkg/framework/safe"
	"context"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
)

// Loader load function signature
type Loader func(context.Context) error

// Ticker handle ticker processes
type Ticker interface {
	Start(context.Context) context.Context
	Notify() <-chan time.Time
}

// NewTicker create new ticker object
func NewTicker(loader Loader, exp, fail time.Duration, retry int) Ticker {
	// NewPool return a new pool object, must start it and watch for ending context
	return &ticker{
		loader:       loader,
		exp:          exp,
		failDuration: fail,
		retry:        retry,
		notify:       make(chan time.Time, 10),
	}
}

type ticker struct {
	loader Loader
	notify chan time.Time
	retry  int

	started      int64
	fail         int
	exp          time.Duration
	failDuration time.Duration
}

func (a *ticker) Start(ctx context.Context) context.Context {
	if !atomic.CompareAndSwapInt64(&a.started, 0, 1) {
		logrus.Panic("already started")
	}
	return safe.ContinuesGoRoutine(ctx, func(cnl context.CancelFunc) time.Duration {
		err := a.loader(ctx)
		if err != nil {
			a.fail++
			if a.fail > a.retry {
				cnl()
				atomic.SwapInt64(&a.started, 0)
			}
			return a.failDuration
		}
		a.fail = 0
		select {
		case a.notify <- time.Now():
		default:
		}
		return a.exp
	})
}

func (a *ticker) Notify() <-chan time.Time {
	return a.notify
}
