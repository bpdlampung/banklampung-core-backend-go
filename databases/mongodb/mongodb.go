package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bpdlampung/banklampung-core-backend-go/entities"
	"github.com/bpdlampung/banklampung-core-backend-go/errors"
	"github.com/bpdlampung/banklampung-core-backend-go/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDBLogger struct {
	mongodb        Mongodb
	collectionName string
}

func NewMongoDBLogger(mongoClient Mongodb, collectionName string) Collections {
	return MongoDBLogger{
		mongodb:        mongoClient,
		collectionName: collectionName,
	}
}

type Find struct {
	Result    interface{}
	CountData *int64
	Filter    interface{}
	Sort      []entities.Sort
	Page      int64
	Size      int64
}

func (f Find) generateOptionSkip() *int64 {
	skipNumber := f.Size * (f.Page - 1)
	return &skipNumber
}

func (m MongoDBLogger) Find(payload Find, ctx context.Context) error {
	start := time.Now()

	collection := m.mongodb.client.Database(m.mongodb.dbname).Collection(m.collectionName)

	findOption := options.Find()

	if payload.Sort != nil && len(payload.Sort) > 0 {
		for x := range payload.Sort {
			findOption.SetSort(bson.D{{payload.Sort[x].FieldName, payload.Sort[x].BuildSortBy()}})
		}
	}

	findOption.Limit = &payload.Size
	findOption.Skip = payload.generateOptionSkip()

	j, _ := json.Marshal(payload.Filter)
	msg := fmt.Sprintf("query: %s", string(j))
	m.mongodb.logger.Debug(msg)

	cursor, err := collection.Find(ctx, payload.Filter, findOption)

	if err != nil {
		msg := fmt.Sprintf("Error Mongodb Connection : %s", err.Error())
		m.mongodb.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, payload.Result); err != nil {
		msg := "cannot unmarshal result"
		m.mongodb.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	finish := time.Now()

	if finish.Sub(start).Seconds() > 10 {
		j, _ := json.Marshal(payload.Filter)
		msg := fmt.Sprintf("slow query: %v second, query: %s", finish.Sub(start).Seconds(), string(j))
		m.mongodb.logger.Debug(msg)
	}

	// handle countdata
	if payload.CountData != nil {
		err := m.Count(CountData{
			Result: payload.CountData,
			Filter: payload.Filter,
		}, ctx)

		if err != nil {
			return err
		}
	}

	marshaledResp, _ := json.Marshal(payload.Result)
	marshaledCount, _ := json.Marshal(payload.CountData)
	m.mongodb.logger.Info(fmt.Sprintf("Success Find Collection: %s, Filter: %s, Result: %s, CountData: %s", m.collectionName, payload.Filter, string(marshaledResp), string(marshaledCount)))

	return nil
}

type FindAll struct {
	Result    interface{}
	CountData *int64
	Filter    interface{}
	Sort      []entities.Sort
}

func (m MongoDBLogger) FindAll(payload FindAll, ctx context.Context) error {
	start := time.Now()

	collection := m.mongodb.client.Database(m.mongodb.dbname).Collection(m.collectionName)

	findOption := options.Find()

	if payload.Sort != nil && len(payload.Sort) > 0 {
		for x := range payload.Sort {
			findOption.SetSort(bson.D{{payload.Sort[x].FieldName, payload.Sort[x].BuildSortBy()}})
		}
	}

	cursor, err := collection.Find(ctx, payload.Filter, findOption)

	if err != nil {
		msg := fmt.Sprintf("Error Mongodb Connection : %s", err.Error())
		m.mongodb.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, payload.Result); err != nil {
		msg := "cannot unmarshal result"
		m.mongodb.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	finish := time.Now()

	if finish.Sub(start).Seconds() > 10 {
		j, _ := json.Marshal(payload.Filter)
		msg := fmt.Sprintf("slow query: %v second, query: %s", finish.Sub(start).Seconds(), string(j))
		m.mongodb.logger.Debug(msg)
	}

	// handle countdata
	if payload.CountData != nil {
		err := m.Count(CountData{
			Result: payload.CountData,
			Filter: payload.Filter,
		}, ctx)

		if err != nil {
			return err
		}
	}

	marshaledResp, _ := json.Marshal(payload.Result)
	marshaledCount, _ := json.Marshal(payload.CountData)
	m.mongodb.logger.Info(fmt.Sprintf("Success Find Collection: %s, Filter: %s, Result: %s, CountData: %s", m.collectionName, payload.Filter, string(marshaledResp), string(marshaledCount)))

	return nil
}

type CountData struct {
	Result *int64
	Filter interface{}
}

func (m MongoDBLogger) Count(payload CountData, ctx context.Context) error {
	start := time.Now()

	collection := m.mongodb.client.Database(m.mongodb.dbname).Collection(m.collectionName)
	countDoc, err := collection.CountDocuments(ctx, payload.Filter)

	if err != nil {
		msg := fmt.Sprintf("Error Mongodb Connection : %s", err.Error())
		m.mongodb.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	if payload.Result != nil {
		*payload.Result = countDoc
	}

	finish := time.Now()

	if finish.Sub(start).Seconds() > 10 {
		j, _ := json.Marshal(payload.Filter)
		msg := fmt.Sprintf("slow query: %v second, query: %s", finish.Sub(start).Seconds(), string(j))
		m.mongodb.logger.Debug(msg)
	}

	return nil
}

type FindOne struct {
	Result interface{}
	Filter interface{}
}

func (m MongoDBLogger) FindOne(payload FindOne, ctx context.Context) error {
	start := time.Now()

	collection := m.mongodb.client.Database(m.mongodb.dbname).Collection(m.collectionName)
	documentReturned := collection.FindOne(ctx, payload.Filter)

	if documentReturned.Err() != nil {
		if documentReturned.Err() == mongo.ErrNoDocuments {
			return nil
		}

		msg := fmt.Sprintf("Error Mongodb Connection : %s", documentReturned.Err())
		m.mongodb.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	if err := documentReturned.Decode(payload.Result); err != nil {
		msg := "cannot unmarshal result"
		m.mongodb.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	finish := time.Now()

	if finish.Sub(start).Seconds() > 10 {
		j, _ := json.Marshal(payload.Filter)
		msg := fmt.Sprintf("slow query: %v second, query: %s", finish.Sub(start).Seconds(), string(j))
		m.mongodb.logger.Debug(msg)
	}

	marshaledResp, _ := json.Marshal(payload.Result)
	m.mongodb.logger.Info(fmt.Sprintf("Success FindOne Collection: %s, Filter: %s, Result: %s", m.collectionName, payload.Filter, string(marshaledResp)))

	return nil
}

type InsertOne struct {
	Result   *string
	Document interface{}
}

func (m MongoDBLogger) InsertOne(payload InsertOne, ctx context.Context) error {
	start := time.Now()

	collection := m.mongodb.client.Database(m.mongodb.dbname).Collection(m.collectionName)
	insertDoc, err := collection.InsertOne(ctx, payload.Document)

	if err != nil {
		msg := fmt.Sprintf("Error Mongodb Connection : %s", err.Error())
		m.mongodb.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	if payload.Result != nil {
		audit := &entities.Audit{}
		helpers.InterfaceToStruct(payload.Document, audit)

		if len(audit.Id) != 0 {
			*payload.Result = insertDoc.InsertedID.(string)
		} else {
			*payload.Result = insertDoc.InsertedID.(primitive.ObjectID).Hex()
		}
	}

	finish := time.Now()

	if finish.Sub(start).Seconds() > 10 {
		j, _ := json.Marshal(payload)
		msg := fmt.Sprintf("slow query: %v second, query: %s", finish.Sub(start).Seconds(), string(j))
		m.mongodb.logger.Debug(msg)
	}

	m.mongodb.logger.Info(fmt.Sprintf("Success InsertOne Collection: %s, Data: %s", m.collectionName, payload.Document))

	return nil
}

type UpdateOne struct {
	Filter   interface{}
	Document interface{}
}

func (m MongoDBLogger) UpdateOne(payload UpdateOne, ctx context.Context) error {
	start := time.Now()

	collection := m.mongodb.client.Database(m.mongodb.dbname).Collection(m.collectionName)

	pByte, err := bson.Marshal(payload.Document)
	if err != nil {
		msg := fmt.Sprintf("Error Mongodb: %s", err.Error())
		m.mongodb.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	var update bson.M
	err = bson.Unmarshal(pByte, &update)
	if err != nil {
		msg := fmt.Sprintf("Error Mongodb: %s", err.Error())
		m.mongodb.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	doc := bson.D{{Key: "$set", Value: update}}
	_, err = collection.UpdateOne(ctx, payload.Filter, doc)

	if err != nil {
		msg := fmt.Sprintf("Error Mongodb Connection : %s", err.Error())
		m.mongodb.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	finish := time.Now()

	if finish.Sub(start).Seconds() > 10 {
		j, _ := json.Marshal(payload.Filter)
		msg := fmt.Sprintf("slow query: %v second, query: %s", finish.Sub(start).Seconds(), string(j))
		m.mongodb.logger.Debug(msg)
	}

	m.mongodb.logger.Info(fmt.Sprintf("Success UpdateOne Collection: %s, Data: %s", m.collectionName, payload.Document))

	return nil
}

type Aggregate struct {
	Result interface{}
	Filter interface{}
}

func (m MongoDBLogger) Aggregate(payload Aggregate, ctx context.Context) error {
	start := time.Now()

	collection := m.mongodb.client.Database(m.mongodb.dbname).Collection(m.collectionName)

	cursor, err := collection.Aggregate(ctx, payload.Filter)

	if err != nil {
		msg := fmt.Sprintf("Error Mongodb Connection : %s", err.Error())
		m.mongodb.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, payload.Result); err != nil {
		msg := "cannot unmarshal result"
		m.mongodb.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	finish := time.Now()

	if finish.Sub(start).Seconds() > 10 {
		j, _ := json.Marshal(payload.Filter)
		msg := fmt.Sprintf("slow query: %v second, query: %s", finish.Sub(start).Seconds(), string(j))
		m.mongodb.logger.Debug(msg)
	}

	m.mongodb.logger.Info(fmt.Sprintf("Success Aggregate Collection: %s, Data: %s", m.collectionName, payload.Filter))

	return nil
}

type DeleteOne struct {
	Filter interface{}
}

func (m MongoDBLogger) DeleteOne(payload DeleteOne, ctx context.Context) error {
	start := time.Now()

	collection := m.mongodb.client.Database(m.mongodb.dbname).Collection(m.collectionName)

	_, err := collection.DeleteOne(ctx, payload.Filter)

	if err != nil {
		msg := fmt.Sprintf("Error Mongodb Connection : %s", err.Error())
		m.mongodb.logger.Error(msg)
		return errors.InternalServerError(msg)
	}

	finish := time.Now()

	if finish.Sub(start).Seconds() > 10 {
		j, _ := json.Marshal(payload.Filter)
		msg := fmt.Sprintf("slow query: %v second, query: %s", finish.Sub(start).Seconds(), string(j))
		m.mongodb.logger.Debug(msg)
	}

	m.mongodb.logger.Info(fmt.Sprintf("Success DeleteOne Collection: %s, Data: %s", m.collectionName, payload.Filter))

	return nil
}
