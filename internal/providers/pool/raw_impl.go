package pool

import (
	"context"
	"github.com/amirex128/selloora_backend/internal/providers/safe"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
)

type rawPl struct {
	loader       RawLoader
	exp          time.Duration
	failDuration time.Duration
	driver       RawDriver
	retry        int

	started int64
	fail    int

	notify chan time.Time
}

func (p *rawPl) Notify() <-chan time.Time {
	return p.notify
}

func (p *rawPl) All() interface{} {
	return p.driver.All()
}

func (p *rawPl) Start(ctx context.Context) context.Context {
	if !atomic.CompareAndSwapInt64(&p.started, 0, 1) {
		logrus.Panic("already started")
	}
	return safe.ContinuesGoRoutine(ctx, func(cnl context.CancelFunc) time.Duration {
		data, err := p.loader(ctx)
		if err != nil {
			p.fail++
			if p.fail > p.retry {
				cnl()
				atomic.SwapInt64(&p.started, 0)
			}
			return 0
		}
		// There is no need to lock here. implementation decide if it require a lock or not
		err = p.driver.Store(data, time.Duration(p.retry)*p.exp)
		if err != nil {
			p.fail++
			return p.failDuration
		}
		p.fail = 0
		select {
		case p.notify <- time.Now():
		default:
		}
		return p.exp
	})
}

// NewPool return a new pool object, must start it and watch for ending context
func NewRawPool(loader RawLoader, driver RawDriver, exp, fail time.Duration, retry int) RawInterface {
	return &rawPl{
		loader:       loader,
		driver:       driver,
		exp:          exp,
		failDuration: fail,
		retry:        retry,
		notify:       make(chan time.Time, 10),
	}
}
