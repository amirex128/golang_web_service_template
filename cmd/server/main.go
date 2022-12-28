package main

import (
	"context"
	"fmt"
	"github.com/amirex128/selloora_backend/internal/api"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title Selloora Backend API
// @version 1.0
// @description Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjUyNzIyMjA0MDcsImV4cGlyZV9hdCI6IiIsImlkIjoxMDEsImlzX2FkbWluIjpmYWxzZSwib3JpZ19pYXQiOjE2NzIyMjQwMDd9.HH7nQlv7qCWElCQv7KYcKk0qhUulx9RjLi-xTS87Cmg

// @contact.name API Support
// @contact.url https://www.amirshirdel.ir
// @contact.email amirex128@gmail.com

// @host localhost:8585
// @BasePath /
// @schemes http https

func main() {
	ctx, _ := context.WithCancel(context.Background())
	models.Initialize(ctx)

	//global_loaders.Initializer(ctx)
	r := api.Runner()
	err := r.Run(viper.GetString("server_host") + ":" + viper.GetString("server_port"))
	if err != nil {
		panic(err)
	}

	sig := WaitExitSignal()
	fmt.Println(sig.String())
}
func WaitExitSignal() os.Signal {
	quit := make(chan os.Signal, 6)
	signal.Notify(quit, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	return <-quit
}
