package mongodb

import (
	"context"
	"github.com/bpdlampung/banklampung-core-backend-go/helpers/structs"
	"go.mongodb.org/mongo-driver/bson"
)

type Templates interface {
	Save(obj interface{}) error
	Update(obj interface{}) error
	GetById(id string, result interface{}) error
}

type implementTemplateRepository struct {
	mongodb Collections
}

func ImplementTemplateRepository(mongoDb Collections) Templates {
	return implementTemplateRepository{
		mongodb: mongoDb,
	}
}

func (i implementTemplateRepository) Save(obj interface{}) error {
	if len(structs.GetStringValueFromStruct(obj, "bson", "_id")) == 0 {
		panic("_id is required")
	}

	err := i.mongodb.InsertOne(InsertOne{
		Document: structs.DoCreateMongoEntity(obj),
	}, context.Background())

	if err != nil {
		return err
	}

	return nil
}

func (i implementTemplateRepository) Update(obj interface{}) error {
	if len(structs.GetStringValueFromStruct(obj, "bson", "_id")) == 0 {
		panic("_id is required")
	}

	err := i.mongodb.UpdateOne(UpdateOne{
		Filter: bson.M{
			"_id": structs.GetStringValueFromStruct(obj, "bson", "_id"),
		},
		Document: structs.DoUpdateMongoEntity(obj),
	}, context.Background())

	if err != nil {
		return err
	}

	return nil
}

func (i implementTemplateRepository) GetById(id string, result interface{}) error {
	if err := i.mongodb.FindOne(FindOne{
		Result: result,
		Filter: bson.M{
			"_id": id,
		},
	}, context.Background()); err != nil {
		return err
	}

	return nil
}
