package providers

import (
	"context"
	"github.com/amirex128/selloora_backend/internal/providers/safe"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
	"time"
)

type LogrusProvider struct {
	Log *logrus.Logger
}

func NewLogrus(i *do.Injector, ctx context.Context) (*LogrusProvider, error) {
	logrusIns := new(LogrusProvider)
	safe.Try(func() error {
		log := logrus.New()
		elasticProvider, err := do.InvokeNamed[*ElasticProvider](i, "main_elastic")
		if err != nil {
			return err
		}
		hook, err := elogrus.NewAsyncElasticHook(elasticProvider.Conn, "localhost", logrus.DebugLevel, "selloora_logs")
		if err != nil {
			return err
		}
		log.Hooks.Add(hook)
		logrusIns.Log = log
		return nil
	}, 30*time.Second)
	logrus.Infof("logrus service is registered")

	return logrusIns, nil
}
