package main

import (
	"net/http"

	"github.com/DanilaNik/IU5_RIP2023/internal/config"
	"github.com/DanilaNik/IU5_RIP2023/internal/dsn"
	"github.com/DanilaNik/IU5_RIP2023/internal/http-server/handlers"
	"github.com/DanilaNik/IU5_RIP2023/internal/minio"
	"github.com/DanilaNik/IU5_RIP2023/internal/repository"
	auth "github.com/DanilaNik/IU5_RIP2023/internal/service/authorization"
	itemservice "github.com/DanilaNik/IU5_RIP2023/internal/service/itemService"

	"github.com/DanilaNik/IU5_RIP2023/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Склад комплектующих.
//Услуги - список комлектующих для хранения с размером для места
//Заявки - заявки на доставку и отгрузку комплектующих

func main() {
	logger := logrus.New()
	router := gin.Default()
	router.Use(CORSMiddleware())
	cfg, err := config.NewConfig(logger)
	if err != nil {
		logger.Fatalf("Error read configuration file: %s", err)
	}
	sourceDB := dsn.FromEnv()
	db, err := repository.NewRepository(sourceDB, logger)
	if err != nil {
		logger.Fatalf("Error init storage: %s", err)
	}
	auth := auth.NewAuthorizationService(db)
	item := itemservice.NewItemService(db, cfg)
	minoCl := minio.NewMinioClient(cfg)
	handler := handlers.NewHandler(logger, db, auth, item, minoCl)
	application := app.New(cfg, router, logger, handler)
	application.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
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

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		c.Header("Access-Control-Allow-Origin", "*")
// 		c.Header("Access-Control-Allow-Credentials", "true")
// 		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }
