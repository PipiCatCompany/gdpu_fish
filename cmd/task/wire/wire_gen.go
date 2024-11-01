// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go-xianyu/internal/server"
	"go-xianyu/pkg/app"
	"go-xianyu/pkg/log"
)

// Injectors from wire.go:

func NewWire(viperViper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	task := server.NewTask(logger)
	appApp := newApp(task)
	return appApp, func() {
	}, nil
}

// wire.go:

var serverSet = wire.NewSet(server.NewTask)

// build App
func newApp(
	task *server.Task,
) *app.App {
	return app.NewApp(app.WithServer(task), app.WithName("demo-task"))
}
