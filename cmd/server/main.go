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
