package mongodb

import "context"

// Collections is mongodb's collection of function
type Collections interface {
	Find(payload FindAllData, ctx context.Context) error
	Count(payload CountData, ctx context.Context) error
	FindOne(payload FindOne, ctx context.Context) error
	InsertOne(payload InsertOne, ctx context.Context) error
	UpdateOne(payload UpdateOne, ctx context.Context) error
	Aggregate(payload Aggregate, ctx context.Context) error
	FindMany(payload FindMany, ctx context.Context) error
}
