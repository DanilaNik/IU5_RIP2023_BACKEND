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
	a.Router.POST("/items/:id/post", a.Handler.PostItemToOrder)

	// a.Router.GET("/users", a.Handler.GetUsers)
	// a.Router.GET("/user", a.Handler.GetUserById)
	// a.Router.POST("/user/delete", a.Handler.DeleteUser)

	a.Router.GET("/orders", a.Handler.GetRequests)
	a.Router.GET("/order", a.Handler.GetRequestById)
	a.Router.POST("/order/delete", a.Handler.DeleteRequest)

	// admin and moderator handlers
	// a.Router.GET("/user/orders", a.RoleMiddleware(role.Admin, role.Moderator), a.Handler.GetUserRequests)
	// a.Router.POST("item", a.RoleMiddleware(role.Admin, role.Moderator), a.Handler.CreateItem)

	//Authorization
	a.Router.POST("/signup", a.Handler.SignUp)
	a.Router.POST("/login", a.Handler.Login)
	a.Router.POST("/logout", a.Handler.Logout)

	a.Router.GET("/protected", a.RoleMiddleware(role.Admin, role.Moderator), a.Handler.ProtectedTest)

	serverAddress := fmt.Sprintf("%s:%d", a.Config.ServiceHost, a.Config.ServicePort)
	if err := a.Router.Run(serverAddress); err != nil {
		a.Logger.Fatal(err)
	}
	a.Logger.Error("Server down")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
