package minio

import (
	"context"
	"fmt"
	"github.com/bpdlampung/banklampung-core-backend-go/logs"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	logger          logs.Collections
	minioClient     *minio.Client
}

func InitConfig(endpoint string, accessKeyID string, secretAccessKey string, useSSL bool, logger logs.Collections) MinioConfig {

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		logger.Error(fmt.Sprintf("Cannot Init Connection Minio!! Error:: %s", err.Error()))
		panic(err)
	}

	return MinioConfig{
		Endpoint:        endpoint,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		UseSSL:          useSSL,
		logger:          logger,
		minioClient:     minioClient,
	}

}

func (m MinioConfig) FGetObject(bucketName, objectName string, fileLoc string) error {
	m.logger.Info(fmt.Sprintf("Bucket Name %s", bucketName))
	m.logger.Info(fmt.Sprintf("Object Name %s", objectName))
	m.logger.Info(fmt.Sprintf("File Loc %s", fileLoc))

	err := m.minioClient.FGetObject(context.Background(), bucketName, objectName, fileLoc, minio.GetObjectOptions{})

	if err != nil {
		return err
	}

	return nil
}
