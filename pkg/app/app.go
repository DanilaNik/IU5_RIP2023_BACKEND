package app

import (
	"fmt"
	"net/http"

	"github.com/DanilaNik/IU5_RIP2023/internal/config"
	"github.com/DanilaNik/IU5_RIP2023/internal/http-server/handlers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Application struct {
	Config  *config.Config
	Logger  *logrus.Logger
	Router  *gin.Engine
	Handler *handlers.Handler
}

func New(c *config.Config, r *gin.Engine, l *logrus.Logger, h *handlers.Handler) *Application {
	return &Application{
		Config:  c,
		Router:  r,
		Logger:  l,
		Handler: h,
	}
}

func (a *Application) Run() {
	a.Logger.Info("Starting service")

	a.Router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	registerStatic(a.Router)

	a.Router.GET("/items", a.Handler.JSONGetItems)
	a.Router.GET("/item/:id", a.Handler.JSONGetItemById)
	a.Router.POST("/item/delete/:id", a.Handler.JSONDeleteItem)
	a.Router.GET("/users", a.Handler.JSONGetUsers)
	a.Router.GET("/user/:id", a.Handler.JSONGetUserById)
	a.Router.POST("/user/delete/:id", a.Handler.JSONDeleteUser)
	a.Router.GET("/orders", a.Handler.JSONGetRequests)
	a.Router.GET("/order/:id", a.Handler.JSONGetRequestById)
	a.Router.POST("/order/delete/:id", a.Handler.JSONDeleteRequest)
	//a.Router.GET("/orders/user/:id", a.Handler.JSONGetUserRequests)

	serverAddress := fmt.Sprintf("%s:%d", a.Config.ServiceHost, a.Config.ServicePort)
	if err := a.Router.Run(serverAddress); err != nil {
		a.Logger.Fatal(err)
	}
	a.Logger.Error("Server down")
}

//func StartServer() {
//	// cfg := config.MustLoad()
//
//	// log := setupLogger(cfg.Env)
//	// log.Info("starting service", slog.String("env", cfg.Env))
//	// log.Debug("debug messages are enabled")
//
//	r := gin.Default()
//	r.GET("/ping", func(ctx *gin.Context) {
//		ctx.JSON(http.StatusOK, gin.H{
//			"message": "pong",
//		})
//	})
//
//	registerStatic(r)
//
//	r.GET("/items", item.New)
//	r.GET("/item/:id", itemCard.New)
//
//	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//
//	// log.Error("Server down")
//}

//const (
//	envLocal = "local"
//	envDev   = "dev"
//	envProd  = "prod"
//)
//
//func setupLogger(env string) *slog.Logger {
//	var log *slog.Logger
//
//	switch env {
//	case envLocal:
//		log = slog.New(
//			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
//		)
//	case envDev:
//		log = slog.New(
//			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
//		)
//	case envProd:
//		log = slog.New(
//			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
//		)
//	}
//	return log
//}

func registerStatic(router *gin.Engine) {
	router.LoadHTMLGlob("templates/html/*")
	router.Static("/templates", "./templates")
	router.Static("/css", "./templates")
	router.Static("/image", "./resources")
}
