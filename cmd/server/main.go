package main

import (
	"context"
	"fmt"
	"github.com/amirex128/selloora_backend/internal/app/api"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/bredis"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/config"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/container"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/mysql"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/signal"
	"github.com/amirex128/selloora_backend/internal/pkg/framework/xlog"
	"github.com/spf13/viper"
)

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
