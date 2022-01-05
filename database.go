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

//Database interface.
type Database interface {
	Disconnect(ctx context.Context) error
	Repository(db DBName, col Collection) Repository
}

//MonogDB implements database interface.
type MongoDB struct {
	c *mongo.Client
}

//NewMongoDB initializes and returns a new MongoDB instace.
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

//Disconnect disconnects db.
func (db *MongoDB) Disconnect(ctx context.Context) error {
	return db.c.Disconnect(ctx)
}

//Repository returns a new repository.
func (db *MongoDB) Repository(dbName DBName, col Collection) Repository {
	collection := db.c.Database(string(dbName)).Collection(string(col))
	return NewRepository(collection)
}
