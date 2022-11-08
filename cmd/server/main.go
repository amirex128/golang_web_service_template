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
)
import "github.com/spf13/viper"

const (
	appName = "selloora"
)

func main() {
	//err := sentry.Init(sentry.ClientOptions{
	//	// Either set your DSN here or set the SENTRY_DSN environment variable.
	//	Dsn: "https://6fa4045bd9ae4e459a87550f63997177@o257983.ingest.sentry.io/4503915107909632",
	//	// Either set environment and release here or set the SENTRY_ENVIRONMENT
	//	// and SENTRY_RELEASE environment variables.
	//	Environment: "",
	//	Release:     "selloora@1.0.0",
	//	// Enable printing of SDK debug messages.
	//	// Useful when getting started or trying to figure something out.
	//	Debug: true,
	//	// Set TracesSampleRate to 1.0 to capture 100%
	//	// of transactions for performance monitoring.
	//	// We recommend adjusting this value in production,
	//	TracesSampleRate: 1.0,
	//	AttachStacktrace: true,
	//})
	//if err != nil {
	//	log.Fatalf("sentry.Init: %s", err)
	//}
	//// Flush buffered events before the program terminates.
	//// Set the timeout to the maximum duration the program can afford to wait.
	//defer sentry.Flush(2 * time.Second)
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
