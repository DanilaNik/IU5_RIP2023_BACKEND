package app

import (
	"net/http"
	"os"

	"github.com/DanilaNik/IU5_RIP2023/internal/http-server/handlers/item"
	"github.com/DanilaNik/IU5_RIP2023/internal/http-server/handlers/itemCard"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

func StartServer() {
	// cfg := config.MustLoad()

	// log := setupLogger(cfg.Env)
	// log.Info("starting service", slog.String("env", cfg.Env))
	// log.Debug("debug messages are enabled")

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	registerStatic(r)

	r.GET("/items", item.New)
	r.GET("/item/:id", itemCard.New)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// log.Error("Server down")
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func registerStatic(router *gin.Engine) {
	router.LoadHTMLGlob("templates/html/*")
	router.Static("/templates", "./templates")
	router.Static("/css", "./templates")
	router.Static("/image", "./resources")
}
