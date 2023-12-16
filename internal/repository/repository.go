package repository

import (
	"os"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	logger *logrus.Logger
	db     *gorm.DB
	rd     *redis.Client
}

func NewRepository(dsn string, log *logrus.Logger) (*Repository, error) {
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ENDPOINT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return &Repository{
		db:     gormDB,
		rd:     redisClient,
		logger: log,
	}, nil
}
