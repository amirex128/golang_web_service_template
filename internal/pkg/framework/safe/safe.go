package safe

import (
	"github.com/amirex128/selloora_backend/internal/pkg/framework/xlog"
	"context"
	"runtime/debug"
	"sync"
	"time"

	"errors"
	"github.com/sirupsen/logrus"
)

type RecoverHook interface {
	Recover(error, []byte, ...interface{})
}

var (
	recoverHooks []RecoverHook
	lock         = &sync.RWMutex{}
)

// Register is a way to register a hook to trigger after panic
func Register(hook RecoverHook) {
	lock.Lock()
	defer lock.Unlock()

	recoverHooks = append(recoverHooks, hook)
}

func Try(a func() error, max time.Duration, extra ...interface{}) {
	x, y := 0, 1
	for {
		err := actual(a, extra...)
		if err == nil {
			return
		}
		t := time.Duration(x) * time.Second
		if t < max {
			x, y = y, x+y
		}
		time.Sleep(t)
	}
}

func actual(a func() error, extra ...interface{}) (e error) {
	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		stack := debug.Stack()
	//		xlog.GetWithError(context.Background(), errors.New("actual failed")).Error(string(stack))
	//		title := mkTitle(err)
	//		call(errors.New(title), stack, extra...)
	//	}
	//}()
	e = a()
	return
}

func Routine(f func(), extra ...interface{}) {
	defer func() {
		e := recover()
		if e != nil {
			stack := debug.Stack()
			xlog.GetWithError(context.Background(), errors.New("Routine failed")).Error(string(stack))
			title := mkTitle(e)
			call(errors.New(title), stack, extra...)
		}
	}()
	f()
}

func call(err error, stack []byte, extra ...interface{}) {
	newExtra := make([]interface{}, 0)
	// if there is an function the call it here as call back.
	// not a cool idea but leave it here for now.
	for i := range extra {
		if fn, ok := extra[i].(func()); ok {
			fn()
		} else {
			newExtra = append(newExtra, extra[i])
		}
	}
	go func() {
		lock.RLock()
		defer lock.RUnlock()
		defer func() {
			if e := recover(); e != nil {
				logrus.Error("What? the recover function is panicked!")
				logrus.Error(e)
			}
		}()

		for i := range recoverHooks {
			recoverHooks[i].Recover(err, stack, newExtra...)
		}
	}()
}

// ContinuesGoRoutine is a safe go routine system with recovery, its continue after recovery
func ContinuesGoRoutine(c context.Context, f func(x context.CancelFunc) time.Duration, extra ...interface{}) context.Context {
	ctx, cl := context.WithCancel(c)
	var s time.Duration
	go func() {
		for i := 1; ; i++ {
			Routine(func() { s = f(cl) }, extra...)
			select {
			case <-ctx.Done():
				logrus.Debug("finalize function and exit")
				return
			case <-time.After(s):
				logrus.Debugf("restart the routine for %d time", i)
			}
		}
	}()
	return ctx
}

// GoRoutine is a safe go routine system with recovery and a way to inform finish of the routine
func GoRoutine(c context.Context, f func(), extra ...interface{}) context.Context {
	ctx, cl := context.WithCancel(c)
	go func() {
		defer cl()
		defer func() {
			if e := recover(); e != nil {
				stack := debug.Stack()
				xlog.GetWithError(context.Background(), errors.New("Routine failed")).Error(e)
				xlog.GetWithError(context.Background(), errors.New("Routine failed")).Error(string(stack))
				title := mkTitle(e)
				call(errors.New(title), stack, extra...)
			}
		}()

		f()
	}()

	return ctx
}

func mkTitle(err interface{}) string {
	var title string
	switch err.(type) {
	case string:
		title = err.(string)
	case error:
		title = err.(error).Error()
	case *logrus.Entry:
		title = err.(*logrus.Entry).Message
	}
	return title
}
