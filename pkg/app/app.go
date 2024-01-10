package app

import (
	"fmt"
	"net/http"

	"github.com/DanilaNik/IU5_RIP2023/docs"
	"github.com/DanilaNik/IU5_RIP2023/internal/config"
	"github.com/DanilaNik/IU5_RIP2023/internal/http-server/handlers"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	docs.SwaggerInfo.Title = "Warehouse"
	docs.SwaggerInfo.Description = "API SERVER"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "172.20.10.6:7070"
	docs.SwaggerInfo.BasePath = "/"

	a.Router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	a.Router.POST("/items/image", a.Handler.LoadS3)

	a.Router.GET("/items", a.Handler.GetItems)
	a.Router.GET("/items/:id", a.Handler.GetItemById)
	a.Router.POST("/items/post", a.Handler.PostItem)
	a.Router.DELETE("/items/:id/delete", a.Handler.DeleteItem)
	a.Router.PUT("/items/:id/put", a.Handler.PutItem)
	a.Router.POST("/items/:id/post", a.Handler.PostItemToRequest)

	a.Router.GET("/orders", a.Handler.GetRequests)
	a.Router.GET("/orders/:id", a.Handler.GetRequestById)
	a.Router.PUT("/orders/:id/approve", a.Handler.PutRequestStatus)
	a.Router.PUT("/orders/make", a.Handler.ConfirmRequest)
	a.Router.DELETE("/order/delete", a.Handler.DeleteRequest)
	a.Router.DELETE("orders/items/:id", a.Handler.DeleteItemFromRequest)

	a.Router.POST("/signup", a.Handler.SignUp)
	a.Router.POST("/login", a.Handler.Login)
	a.Router.POST("/logout", a.Handler.Logout)

	a.Router.Use(a.Handler.UserAuth).POST("/validate", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	})
	a.Router.Use(a.Handler.AdminAuth).POST("/validate_admin", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{})
	})

	serverAddress := fmt.Sprintf("%s:%d", a.Config.ServiceHost, a.Config.ServicePort)
	if err := a.Router.Run(serverAddress); err != nil {
		a.Logger.Fatal(err)
	}
	a.Logger.Error("Server down")
}
