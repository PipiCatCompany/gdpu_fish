//go:build wireinject
// +build wireinject

package wire

import (
	"go-xianyu/internal/handler"
	"go-xianyu/internal/repository"
	"go-xianyu/internal/server"
	"go-xianyu/internal/service"
	"go-xianyu/pkg/app"
	"go-xianyu/pkg/jwt"
	"go-xianyu/pkg/log"
	"go-xianyu/pkg/server/http"
	"go-xianyu/pkg/sid"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewPostRepository,
	repository.NewCommentRepository,
	repository.NewMessageRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewPostService,
	service.NewCommentService,
	service.NewMessageService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewPostHandler,
	handler.NewCommentHandler,
	handler.NewMessageHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
)

// build App
func newApp(
	httpServer *http.Server,
	job *server.Job,
	// task *server.Task,
) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
	))
}
