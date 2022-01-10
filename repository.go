package dataful

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Repository interface.
type Repository interface {
	InsertOne(ctx context.Context, document Model, opts ...*options.InsertOneOptions) error
	FindOne(ctx context.Context, filter bson.M, document Model, opts ...*options.FindOneOptions) error
	ReplaceOne(ctx context.Context, filter bson.M, replacement Model, opts ...*options.ReplaceOptions) error
	FindAll(ctx context.Context, filter bson.M, list ModelList, opts ...*options.FindOptions) error
	DeleteOne(ctx context.Context, filter bson.M, opts ...*options.DeleteOptions) error
}

//MongoRepository implements the repository interface.
type MongoRepository struct {
	c *mongo.Collection
}

//NewRepository returns a new mongo repository.
func NewRepository(c *mongo.Collection) *MongoRepository {
	return &MongoRepository{
		c: c,
	}
}

//FindOne finds one document from db.
func (r *MongoRepository) FindOne(ctx context.Context, filter bson.M, document Model, opts ...*options.FindOneOptions) error {
	err := r.c.FindOne(ctx, filter, opts...).Decode(document)
	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		return err
	}
	document.SetLoaded(true)
	return nil
}

//FindAll finds all documents
func (r *MongoRepository) FindAll(ctx context.Context, filter bson.M, list ModelList, opts ...*options.FindOptions) error {
	cursor, err := r.c.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	err = cursor.All(ctx, list.GetList())
	if err != nil {
		return err
	}
	list.SetLoaded()
	return nil
}

//ReplaceOne replaced one document in db.
func (r *MongoRepository) ReplaceOne(ctx context.Context, filter bson.M, replacement Model, opts ...*options.ReplaceOptions) error {
	replacement.Updated()
	updateResult, err := r.c.ReplaceOne(ctx, filter, replacement, opts...)
	if err != nil {
		return err
	}
	if updateResult.MatchedCount == 0 {
		return errors.New("zero matches returned")
	}
	return nil
}

//InsertOne inserts one model to db.
func (r *MongoRepository) InsertOne(ctx context.Context, document Model, opts ...*options.InsertOneOptions) error {
	document.Created()
	insertResult, err := r.c.InsertOne(ctx, document, opts...)
	if err != nil {
		return err
	}

	document.SetID(insertResult.InsertedID)
	document.SetLoaded(true)
	return nil
}

//DeleteOne deletes one document on db.
func (r *MongoRepository) DeleteOne(ctx context.Context, filter bson.M, opts ...*options.DeleteOptions) error {
	result, err := r.c.DeleteOne(ctx, filter, opts...)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("zero matches returned")
	}
	return nil
}
