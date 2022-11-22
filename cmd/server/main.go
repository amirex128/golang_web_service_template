package main

import (
	"context"
	"fmt"
	"github.com/amirex128/selloora_backend/internal/api"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.elastic.co/apm/module/apmlogrus"
	"os"
	"os/signal"
	"syscall"
)

const (
	appName = "selloora"
)

// @title Selloora Backend API
// @version 1.0
// @description Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFtaXJleDEyOEBnbWFpbC5jb20iLCJleHAiOjUyNjgzMTg5OTYsImV4cGlyZV9hdCI6IiIsImZpcnN0bmFtZSI6Itin2YXbjNixIiwiaWQiOjEsImxhc3RuYW1lIjoi2LTbjNix2K_ZhNuMIiwibW9iaWxlIjoiMDkwMjQ4MDk3NTAiLCJvcmlnX2lhdCI6MTY2ODMyMjU5Niwic3RhdHVzIjoiIn0.x7BKuxw288cm1JsskGRD178UPmNz-xRwkWHtb0WsU74

// @contact.name API Support
// @contact.url https://www.amirshirdel.ir
// @contact.email amirex128@gmail.com

// @host localhost:8585
// @BasePath /
// @schemes http https

func main() {
	ctx, _ := context.WithCancel(context.Background())

	InitializeConfig(appName)
	models.Initialize(ctx)

	//global_loaders.Initializer(ctx)
	api.Runner(viper.GetString("server_host"), viper.GetString("server_port"), ctx)

	sig := WaitExitSignal()
	fmt.Println(sig.String())
}
func WaitExitSignal() os.Signal {
	quit := make(chan os.Signal, 6)
	signal.Notify(quit, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	return <-quit
}

func InitializeConfig(prefix string) {
	logrus.AddHook(&apmlogrus.Hook{})
	logrus.SetOutput(os.Stdout)
	viper.AutomaticEnv()

}
