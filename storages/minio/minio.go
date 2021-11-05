package minio

import (
	"bytes"
	"context"
	"fmt"
	"github.com/bpdlampung/banklampung-core-backend-go/errors"
	"github.com/minio/minio-go/v7"
	"io/ioutil"
	"os"
	"time"
)

type Put struct {
	Context      context.Context
	BucketName   string
	FileLocation string
	FileName     string
}

func (m Minio) Put(put Put) (*string, error) {
	ok, err := m.client.BucketExists(put.Context, put.BucketName)

	if err != nil {
		return nil, errors.InternalServerError(fmt.Sprintf("Cannot do Bucket Exists Check ::%s", err.Error()))
	}

	if !ok {
		err = m.client.MakeBucket(put.Context, put.BucketName, minio.MakeBucketOptions{})

		if err != nil {
			return nil, errors.InternalServerError("Cannot do Make Bucket")
		}
	}

	finput, err := os.Open(put.FileLocation)

	if err != nil {
		return nil, errors.BadRequest(fmt.Sprintf("File not found ::%s", put.FileLocation))
	}

	input, err := ioutil.ReadAll(finput)

	if err != nil {
		return nil, err
	}

	in := bytes.NewReader(input)

	defer finput.Close()

	_, err = m.client.PutObject(put.Context, put.BucketName, put.FileName, in, in.Size(), minio.PutObjectOptions{})

	if err != nil {
		return nil, err
	}

	return m.Get(Get{
		Context:    put.Context,
		BucketName: put.BucketName,
		FileName:   put.FileName,
		Expires:    5 * time.Minute,
	})
}

type Get struct {
	Context    context.Context
	BucketName string
	FileName   string
	Expires    time.Duration
}

func (m Minio) Get(get Get) (*string, error) {
	url, err := m.client.PresignedGetObject(context.Background(), get.BucketName, get.FileName, get.Expires, nil)

	if err != nil {
		return nil, errors.InternalServerError(fmt.Sprintf("Cannot get object ::%s", err.Error()))
	}

	urlString := url.String()

	return &urlString, nil
}