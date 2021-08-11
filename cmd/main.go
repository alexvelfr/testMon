package main

import (
	"context"

	"github.com/alexvelfr/tmp/server"
	"github.com/spf13/viper"
)

func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()

}

func main() {
	if err := initConfig(); err != nil {
		panic("config not set")
	}
	ctx, cancle := context.WithCancel(context.Background())
	app := server.NewApp(ctx, cancle)
	app.Run()
	<-ctx.Done()
}
