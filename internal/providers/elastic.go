package providers

import (
	"context"
	"fmt"
	"github.com/amirex128/selloora_backend/internal/providers/safe"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.elastic.co/apm/module/apmelasticsearch/v2"
	"net/http"
	"time"
)
import "github.com/olivere/elastic/v7"

type ElasticProvider struct {
	Conn *elastic.Client
	Ctx  context.Context
}

func NewElastic(i *do.Injector, ctx context.Context) (*ElasticProvider, error) {

	elasticIns := new(ElasticProvider)

	safe.Try(func() error {

		connection := fmt.Sprintf("http://%s:%s", viper.GetString("SELLOORA_ELASTIC_HOST"), viper.GetString("SELLOORA_ELASTIC_PORT"))

		var err error
		elasticIns.Conn, err = elastic.NewClient(
			elastic.SetURL(connection),
			elastic.SetSniff(false),
			elastic.SetHealthcheck(false),
			elastic.SetHttpClient(&http.Client{
				Transport: apmelasticsearch.WrapRoundTripper(http.DefaultTransport),
			}),
		)
		if err != nil {
			logrus.Errorf("failed to register elastic service")
			return err
		}
		return nil
	}, 30*time.Second)
	logrus.Infof("elastic service is registered")
	return elasticIns, nil
}
