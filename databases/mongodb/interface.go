package mongodb

import "context"

// Collections is mongodb's collection of function
type Collections interface {
	Find(payload Find, ctx context.Context) error
	FindAll(payload FindAll, ctx context.Context) error // without page and size
	Count(payload CountData, ctx context.Context) error
	FindOne(payload FindOne, ctx context.Context) error
	InsertOne(payload InsertOne, ctx context.Context) error
	UpdateOne(payload UpdateOne, ctx context.Context) error
	Aggregate(payload Aggregate, ctx context.Context) error
	DeleteOne(payload DeleteOne, ctx context.Context) error
}
