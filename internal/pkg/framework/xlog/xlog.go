package xlog

import (
	"backend/internal/pkg/framework/assert"
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.elastic.co/apm/module/apmlogrus/v2"
	"net"
	"os"
	"sync"
	"time"
)

type contextKey int

const ctxKey contextKey = iota

var loc *time.Location

var specialLogger *logrus.Logger

type concurrentFields struct {
	fields logrus.Fields
	lock   sync.RWMutex
}

func GetSpeciallog() *logrus.Logger {
	return specialLogger
}

type TehranFormatter struct {
	logrus.Formatter
}

func (u TehranFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.In(loc)
	return u.Formatter.Format(e)
}

func Initialize(appName string) {
	logrus.SetFormatter(TehranFormatter{&logrus.JSONFormatter{}})
	if viper.GetBool("develop_mode") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	if viper.GetBool("stash_log") {
		addStashHook(appName)
	}

	if viper.GetBool("log_to_file") {
		if viper.GetBool("log_info") {
			setupFileWriterHook(appName, "info", []logrus.Level{logrus.InfoLevel}, TehranFormatter{&logrus.JSONFormatter{}})
		}
		if viper.GetBool("log_warn") {
			setupFileWriterHook(appName, "warn", []logrus.Level{logrus.WarnLevel}, TehranFormatter{&logrus.JSONFormatter{}})
		}
		setupFileWriterHook(appName, "error", []logrus.Level{logrus.ErrorLevel}, TehranFormatter{&logrus.JSONFormatter{}})
		setupFileWriterHook(appName, "fatal", []logrus.Level{logrus.FatalLevel}, TehranFormatter{&logrus.JSONFormatter{}})
		setupFileWriterHook(appName, "panic", []logrus.Level{logrus.PanicLevel}, TehranFormatter{&logrus.JSONFormatter{}})
	}
	if viper.GetBool("special_log_to_file") {
		setupFileWriterHook(appName, "special", []logrus.Level{
			logrus.DebugLevel,
			logrus.InfoLevel,
			logrus.WarnLevel,
			logrus.ErrorLevel,
			logrus.FatalLevel,
			logrus.PanicLevel},
			TehranFormatter{&logrus.JSONFormatter{}})
	}
	logrus.Info("log initialized")
}

// setupFileWriterHook create a hook with input params and adds it to logrus
// also, every day it will refresh file writer and creat new file and send logs to that.
func setupFileWriterHook(appName string, level string, logLevel []logrus.Level, formatter logrus.Formatter) {
	t := time.Now()
	logPath := viper.GetString("log_path")
	assert.Nil(os.MkdirAll(fmt.Sprintf("%s/%s", logPath, appName), 0755))
	file, err := os.OpenFile(fmt.Sprintf("%s/%s/%s_%d-%d-%d", logPath, appName, level, t.Year(), t.Month(), t.Day()), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	assert.Nil(err)

	hook := &LogHook{
		Writer:    file,
		LogLevels: logLevel,
		Formatter: formatter,
	}
	logrus.AddHook(hook)
	logrus.AddHook(&apmlogrus.Hook{})

	go watchFile(appName, level, hook)
}

func addStashHook(appName string) {
	conn, err := net.Dial("tcp",
		fmt.Sprintf("%s:%d", viper.GetString("stash_host"), viper.GetInt("stash_port")))
	if err != nil {
		logrus.Error(err)
		return
	}

	//hook := New(conn, DefaultFormatter(logrus.Fields{"type": appName}))
	//logrus.AddHook(hook)

	h1 := &LogHook{
		conn,
		DefaultFormatter(logrus.Fields{"type": appName, "level": "warn"}),
		[]logrus.Level{logrus.WarnLevel},
	}
	h2 := &LogHook{
		conn,
		DefaultFormatter(logrus.Fields{"type": appName, "level": "info"}),
		[]logrus.Level{logrus.InfoLevel},
	}

	h3 := &LogHook{
		conn,
		DefaultFormatter(logrus.Fields{"type": appName, "level": "error"}),
		[]logrus.Level{logrus.ErrorLevel},
	}
	h4 := &LogHook{
		conn,
		DefaultFormatter(logrus.Fields{"type": appName, "level": "fatal"}),
		[]logrus.Level{logrus.FatalLevel},
	}
	h5 := &LogHook{
		conn,
		DefaultFormatter(logrus.Fields{"type": appName, "level": "panic"}),
		[]logrus.Level{logrus.PanicLevel},
	}

	if viper.GetBool("log_warn") {
		logrus.AddHook(h1)
	}
	if viper.GetBool("log_info") {
		logrus.AddHook(h2)
	}
	logrus.AddHook(h3)
	logrus.AddHook(h4)
	logrus.AddHook(h5)

	logrus.Warn("warn level test")
	logrus.Error("error level test")
}

// watchFile check for files that logs write into and every day creat new file and replace
// writer in logrus.
func watchFile(appName string, level string, hook *LogHook) {
	t := time.Now()
	logPath := viper.GetString("log_path")
	lastFileName := fmt.Sprintf("%s/%s/%s_%d-%d-%d", logPath, appName, level, t.Year(), t.Month(), t.Day())
	for {
		select {
		case <-time.After(3 * time.Second):
			t = time.Now()
			fName := fmt.Sprintf("%s/%s/%s_%d-%d-%d", logPath, appName, level, t.Year(), t.Month(), t.Day())
			if fName != lastFileName {
				f, err := os.OpenFile(fName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
				assert.Nil(err)
				// Replace new file write
				hook.Writer = f
			}
			lastFileName = fName
		}
	}
}

// Get generate log's entry
func Get(ctx context.Context) *logrus.Entry {
	fields, ok := ctx.Value(ctxKey).(*concurrentFields)
	entry := logrus.NewEntry(logrus.StandardLogger())
	if ok {
		fields.lock.RLock()
		defer fields.lock.RUnlock()
		return entry.WithFields(fields.fields)
	}
	return entry
}

// GetWithError is a shorthand for Get(ctx).WithError(err)
func GetWithError(ctx context.Context, err error) *logrus.Entry {
	return Get(ctx).WithError(err)
}

// GetWithField is a shorthand for Get().WithField()
func GetWithField(ctx context.Context, key string, val interface{}) *logrus.Entry {
	return Get(ctx).WithField(key, val)
}

// GetWithFields is a shorthand for Get().WithFields()
func GetWithFields(ctx context.Context, f logrus.Fields) *logrus.Entry {
	return Get(ctx).WithFields(f)
}

// SetField in the context
func SetField(ctx context.Context, key string, val interface{}) context.Context {
	fields, ok := ctx.Value(ctxKey).(*concurrentFields)
	if !ok {
		fields = &concurrentFields{
			fields: make(logrus.Fields),
			lock:   sync.RWMutex{},
		}
	}
	fields.lock.Lock()
	defer fields.lock.Unlock()
	fields.fields[key] = val

	return context.WithValue(ctx, ctxKey, fields)
}

// SetFields set the fields for logger
func SetFields(ctx context.Context, fl logrus.Fields) context.Context {
	fields, ok := ctx.Value(ctxKey).(*concurrentFields)
	if !ok {
		fields = &concurrentFields{
			fields: make(logrus.Fields),
			lock:   sync.RWMutex{},
		}
	}
	fields.lock.Lock()
	defer fields.lock.Unlock()
	for i := range fl {
		fields.fields[i] = fl[i]
	}
	return context.WithValue(ctx, ctxKey, fields)
}

func init() {
	loc, _ = time.LoadLocation("Asia/Tehran")
}
