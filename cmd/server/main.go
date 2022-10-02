package main

import (
	"backend/api"
	"backend/internal/pkg/framework/bredis"
	"backend/internal/pkg/framework/config"
	"backend/internal/pkg/framework/container"
	"backend/internal/pkg/framework/mysql"
	"backend/internal/pkg/framework/signal"
	"backend/internal/pkg/framework/xlog"
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	"log"
	"time"
)
import "github.com/spf13/viper"

const (
	appName = "selloora"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://6fa4045bd9ae4e459a87550f63997177@o257983.ingest.sentry.io/4503915107909632",
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage("It works!")

	config.Initialize(appName)
	xlog.Initialize(appName)
	container.RegisterServices(appName)
	ctx, _ := context.WithCancel(context.Background())

	bredis.Initialize(ctx)
	mysql.Initialize(ctx)
	//global_loaders.Initializer(ctx)

	api.Runner(viper.GetString("server_host"), viper.GetString("server_port"))

	sig := signal.WaitExitSignal()
	fmt.Println(sig.String())
}
