package mongodb

import (
	"banklampung-core/logs"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type Mongodb struct {
	client *mongo.Client
	dbname string
	logger logs.Collections
}

var mongoClient Mongodb

func GenerateUri(host, port, name, username, password string) string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", username, password, host, port, name)
}

func InitConnection(masterDBUrl string, logger logs.Collections) Mongodb {
	mClient, err := newClient(masterDBUrl)

	if err != nil {
		panic(err)
	}

	mongoClient = Mongodb{
		client: mClient,
		dbname: strings.Split(strings.ReplaceAll(masterDBUrl, "//", ""), "/")[1],
		logger: logger,
	}

	logger.Info("MongoDB Connected")

	return mongoClient
}

func newClient(mongoUri string) (*mongo.Client, error) {
	client, err := mongo.Connect(
		context.Background(),
		options.Client().SetRetryWrites(true),
		options.Client().SetRetryReads(true),
		options.Client().SetMaxPoolSize(1000),
		options.Client().ApplyURI(mongoUri),
	)

	if err != nil {
		return nil, err
	}

	//if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
	//	return nil, err
	//}

	return client, nil
}
