package main

import (
	"github.com/DanilaNik/IU5_RIP2023/internal/config"
	"github.com/DanilaNik/IU5_RIP2023/internal/dsn"
	"github.com/DanilaNik/IU5_RIP2023/internal/http-server/handlers"
	"github.com/DanilaNik/IU5_RIP2023/internal/repository"
	auth "github.com/DanilaNik/IU5_RIP2023/internal/service/authorization"
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
	handler := handlers.NewHandler(logger, db, auth)
	application := app.New(cfg, router, logger, handler)
	application.Run()
}
