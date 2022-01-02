package dataful

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBName string

const (
	DBAmazeful DBName = "Amazeful"
	DBLogs     DBName = "Logs"
	DBTest     DBName = "Test"
)

type Database interface {
	Disconnect(ctx context.Context) error
	Repository(db DBName, col Collection) Repository
}

type MongoDB struct {
	c *mongo.Client
}

func NewMongoDB(ctx context.Context, uri string) (*MongoDB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		c: client,
	}, nil
}

func (db *MongoDB) Disconnect(ctx context.Context) error {
	return db.c.Disconnect(ctx)
}

func (db *MongoDB) Repository(dbName DBName, col Collection) Repository {
	collection := db.c.Database(string(dbName)).Collection(string(col))
	return NewRepository(collection)
}
