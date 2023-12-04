package minio

import (
	"log"

	"github.com/DanilaNik/IU5_RIP2023/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	*minio.Client
}

func NewMinioClient(cfg *config.Config) *MinioClient {

	useSSL := false

	minioClient, err := minio.New(config.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Minio.User, config.Minio.Pass, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return &MinioClient{
		minioClient,
	}
}
