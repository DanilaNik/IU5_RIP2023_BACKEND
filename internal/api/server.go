package api

//
//import (
//	"github.com/DanilaNik/IU5_RIP2023/internal/storage"
//	"github.com/gin-gonic/gin"
//	"golang.org/x/exp/slog"
//	"net/http"
//	"os"
//	"strconv"
//)
//
//func StartServer1() {
//	//cfg := config.MustLoad()
//	//
//	//log := setupLogger(cfg.Env)
//	//log.Info("starting service", slog.String("env", cfg.Env))
//	//log.Debug("debug messages are enabled")
//
//	r := gin.Default()
//	r.GET("/ping", func(ctx *gin.Context) {
//		ctx.JSON(http.StatusOK, gin.H{
//			"message": "pong",
//		})
//	})
//
//	r.LoadHTMLGlob("templates/html/*")
//
//	//r.GET("/items", item.New)
//
//	r.GET("/items", func(ctx *gin.Context) {
//		filter := ctx.Query("filter")
//		quantityMin := ctx.Query("quantitymin")
//		quantityMax := ctx.Query("quantitymax")
//		data := storage.GetItems()
//
//		if filter == "" {
//			filter = "all"
//		}
//		var q1, q2 uint64
//		if quantityMin == "" {
//			q1 = 0
//		} else {
//			q1, err := strconv.ParseUint(quantityMin, 10, 64)
//			_ = q1
//			if err != nil {
//				return
//			}
//		}
//		if quantityMax == "" {
//			q2 = uint64(^uint32(0))
//		} else {
//			q2, err := strconv.ParseUint(quantityMax, 10, 64)
//			_ = q2
//			if err != nil {
//				return
//			}
//		}
//
//		res := make([]*storage.Item, 0)
//
//		for _, v := range data {
//			q := v.Quantity
//			if q < q1 || q >= q2 {
//				continue
//			}
//			if filter == "all" {
//				res = append(res, v)
//			} else if v.Status == filter {
//				res = append(res, v)
//			}
//		}
//
//		ctx.HTML(http.StatusOK, "items.tmpl", data)
//	})
//	registerStatic(r)
//	//r.Static("/image", "./resources")
//
//	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
//
//	//log.Error("Server down")
//}
//
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
//
//func registerStatic(router *gin.Engine) {
//	router.LoadHTMLGlob("templates/html/*")
//	router.Static("/templates", "./templates")
//	router.Static("/css", "./templates")
//	router.Static("/image", "./resources")
//}
