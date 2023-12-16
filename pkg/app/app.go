package app

import (
	"fmt"
	"net/http"

	"github.com/DanilaNik/IU5_RIP2023/docs"
	"github.com/DanilaNik/IU5_RIP2023/internal/config"
	"github.com/DanilaNik/IU5_RIP2023/internal/http-server/handlers"
	"github.com/DanilaNik/IU5_RIP2023/internal/service/role"
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
	docs.SwaggerInfo.Host = "localhost:7070"
	docs.SwaggerInfo.BasePath = "/"

	a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	a.Router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	registerStatic(a.Router)
	a.Router.GET("/items", a.Handler.GetItems)
	a.Router.GET("/item/:id", a.Handler.GetItemById)
	a.Router.POST("/item/delete/:id", a.Handler.DeleteItem)
	a.Router.GET("/users", a.Handler.GetUsers)
	a.Router.GET("/user", a.Handler.GetUserById)
	a.Router.POST("/user/delete", a.Handler.DeleteUser)
	a.Router.GET("/orders", a.Handler.GetRequests)
	a.Router.GET("/order", a.Handler.GetRequestById)
	a.Router.POST("/order/delete", a.Handler.DeleteRequest)

	// admin and moderator handlers
	a.Router.GET("/user/orders", a.Handler.GetUserRequests)

	//Authorization
	a.Router.POST("/register", a.Handler.Registr)
	a.Router.POST("/login", a.Handler.Login)

	a.Router.GET("/protected", a.RoleMiddleware(role.Admin, role.Moderator), a.Handler.ProtectedTest)

	serverAddress := fmt.Sprintf("%s:%d", a.Config.ServiceHost, a.Config.ServicePort)
	if err := a.Router.Run(serverAddress); err != nil {
		a.Logger.Fatal(err)
	}
	a.Logger.Error("Server down")
}

func registerStatic(router *gin.Engine) {
	router.LoadHTMLGlob("templates/html/*")
	router.Static("/templates", "./templates")
	router.Static("/css", "./templates")
	router.Static("/image", "./resources")
}
