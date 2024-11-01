// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go-xianyu/internal/repository"
	"go-xianyu/internal/server"
	"go-xianyu/pkg/app"
	"go-xianyu/pkg/log"
)

// Injectors from wire.go:

func NewWire(viperViper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	db := repository.NewDB(viperViper, logger)
	migrate := server.NewMigrate(db, logger)
	appApp := newApp(migrate)
	return appApp, func() {
	}, nil
}

// wire.go:

var repositorySet = wire.NewSet(repository.NewDB, repository.NewRepository, repository.NewUserRepository)

var serverSet = wire.NewSet(server.NewMigrate)

// build App
func newApp(
	migrate *server.Migrate,
) *app.App {
	return app.NewApp(app.WithServer(migrate), app.WithName("demo-migrate"))
}
