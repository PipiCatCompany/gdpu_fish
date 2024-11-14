package server

import (
	apiV1 "go-xianyu/api/v1"
	"go-xianyu/docs"
	"go-xianyu/internal/handler"
	"go-xianyu/internal/middleware"
	"go-xianyu/pkg/jwt"
	"go-xianyu/pkg/log"
	"go-xianyu/pkg/server/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	postHandler *handler.PostHandler,
	commentHandler *handler.CommentHandler,
	messageHandler *handler.MessageHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "Thank you for using nunu!",
		})
	})

	// TODO 对服务分group
	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			// User-router
			noAuthRouter.POST("/register", userHandler.Register)
			noAuthRouter.POST("/login", userHandler.Login)
			noAuthRouter.GET("/login_openid", userHandler.LoginByOpenId)
			noAuthRouter.POST("/user_auto", userHandler.CreateUserBasic)
			noAuthRouter.GET("/openid", userHandler.GetOpenId)
			// noAuthRouter.POST("/openid2login", userHandler.LoginByOpenId)

			// Comment-router
			noAuthRouter.GET("/comments", commentHandler.GetCommentList)

			// Post-router
			noAuthRouter.GET("/posts", postHandler.GetPostListByPage)
		}

		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		{
			noStrictAuthRouter.GET("/user", userHandler.GetProfile)
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)

			// POST-router
			strictAuthRouter.POST("/post", postHandler.CreatePost)

			// Comment-router
			strictAuthRouter.POST("/comment", commentHandler.CreateComment)

			// User-router
			strictAuthRouter.PUT("/user/studentcode", userHandler.UpdateUserStudentCode)

			// Message-router
			strictAuthRouter.POST("/msg", messageHandler.CreateMessage)
			strictAuthRouter.GET("/msgChanel", messageHandler.GetMessageChanelInfo)
			strictAuthRouter.GET("/msgs", messageHandler.GetMessageByPagination)

			// QiNiu
			noStrictAuthRouter.GET("/qiniu/token", postHandler.GetQiNiuToken)
		}
	}

	return s
}
