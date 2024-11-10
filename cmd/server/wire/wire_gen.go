// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go-xianyu/internal/handler"
	"go-xianyu/internal/repository"
	"go-xianyu/internal/server"
	"go-xianyu/internal/service"
	"go-xianyu/pkg/app"
	"go-xianyu/pkg/jwt"
	"go-xianyu/pkg/log"
	"go-xianyu/pkg/server/http"
	"go-xianyu/pkg/sid"
)

// Injectors from wire.go:

func NewWire(viperViper *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	jwtJWT := jwt.NewJwt(viperViper)
	handlerHandler := handler.NewHandler(logger)
	db := repository.NewDB(viperViper, logger)
	repositoryRepository := repository.NewRepository(logger, db)
	transaction := repository.NewTransaction(repositoryRepository)
	sidSid := sid.NewSid()
	serviceService := service.NewService(transaction, logger, sidSid, jwtJWT)
	userRepository := repository.NewUserRepository(repositoryRepository)
	userService := service.NewUserService(serviceService, userRepository)
	userHandler := handler.NewUserHandler(handlerHandler, userService)
	postRepository := repository.NewPostRepository(repositoryRepository)
	postService := service.NewPostService(serviceService, postRepository, userRepository)
	postHandler := handler.NewPostHandler(handlerHandler, postService)
	commentRepository := repository.NewCommentRepository(repositoryRepository, userRepository)
	commentService := service.NewCommentService(serviceService, commentRepository)
	commentHandler := handler.NewCommentHandler(handlerHandler, commentService)
	messageRepository := repository.NewMessageRepository(repositoryRepository)
	messageService := service.NewMessageService(serviceService, messageRepository)
	messageHandler := handler.NewMessageHandler(handlerHandler, messageService)
	httpServer := server.NewHTTPServer(logger, viperViper, jwtJWT, userHandler, postHandler, commentHandler, messageHandler)
	job := server.NewJob(logger)
	appApp := newApp(httpServer, job)
	return appApp, func() {
	}, nil
}

// wire.go:

var repositorySet = wire.NewSet(repository.NewDB, repository.NewRepository, repository.NewTransaction, repository.NewUserRepository, repository.NewPostRepository, repository.NewCommentRepository, repository.NewMessageRepository)

var serviceSet = wire.NewSet(service.NewService, service.NewUserService, service.NewPostService, service.NewCommentService, service.NewMessageService)

var handlerSet = wire.NewSet(handler.NewHandler, handler.NewUserHandler, handler.NewPostHandler, handler.NewCommentHandler, handler.NewMessageHandler)

var serverSet = wire.NewSet(server.NewHTTPServer, server.NewJob)

// build App
func newApp(
	httpServer *http.Server,
	job *server.Job,

) *app.App {
	return app.NewApp(app.WithServer(httpServer, job), app.WithName("demo-server"))
}
