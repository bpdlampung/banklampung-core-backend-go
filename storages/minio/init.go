package minio

import (
	"github.com/bpdlampung/banklampung-core-backend-go/logs"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	client *minio.Client
	logger logs.Collections
}

func InitConnection(endpoint, accessKey, secretKey string, secure bool, logger logs.Collections) Minio {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: secure,
	})

	if err != nil {
		panic(err)
	}

	return Minio{
		client: minioClient,
		logger: logger,
	}
}
