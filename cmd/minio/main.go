package main

import (
	"context"

	"github.com/DanilaNik/IU5_RIP2023/internal/config"
	"github.com/sirupsen/logrus"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	logger := logrus.New()
	cfg, err := config.NewConfig(logger)
	if err != nil {
		logger.Fatalf("Error read configuration file: %s", err)
	}

	ctx := context.Background()
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(cfg.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.User, cfg.Minio.Pass, ""),
		Secure: useSSL,
	})
	if err != nil {
		logger.Fatalln(err)
	}

	// Make a new bucket called testbucket.
	bucketName := "cnc"
	location := "eu-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			logger.Printf("We already own %s\n", bucketName)
		} else {
			logger.Fatalln(err)
		}
	} else {
		logger.Printf("Successfully created %s\n", bucketName)
	}
}
