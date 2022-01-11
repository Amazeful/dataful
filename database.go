package dataful

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
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

	err = createCollectionsAndIndexes(ctx, client)
	if err != nil {
		return nil, err
	}

	return &MongoDB{c: client}, nil
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

//CollectionStr returns collection in string type.
func CollectionStr(collection Collection) string {
	return string(collection)
}

//DbStr returns database name in string type.
func DbStr(database DBName) string {
	return string(database)
}

//createCollectionsAndIndexes creates all collections and ensures indexes are added in.
func createCollectionsAndIndexes(ctx context.Context, c *mongo.Client) error {
	// ==============================|| AMAZEFULDB ||============================== //
	amazefulDB := c.Database(DbStr(DBAmazeful))

	//Alerts
	_, err := amazefulDB.Collection(CollectionStr(CollectionAlerts)).Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.M{"channel": 1}, Options: options.Index().SetUnique(true)},
	})
	if err != nil {
		return err
	}

	//Channel
	_, err = amazefulDB.Collection(CollectionStr(CollectionChannel)).Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.M{"channelId": 1}, Options: options.Index().SetUnique(true)},
		{Keys: bson.M{"broadcasterName": 1}},
		{Keys: bson.M{"joined": 1}},
	})
	if err != nil {
		return err
	}

	//Command
	_, err = amazefulDB.Collection(CollectionStr(CollectionCommand)).Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{"channel", 1}, {"name", 1}}, Options: options.Index().SetUnique(true)},
	})
	if err != nil {
		return err
	}

	//Filters
	_, err = amazefulDB.Collection(CollectionStr(CollectionFilters)).Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.M{"channel": 1}, Options: options.Index().SetUnique(true)},
	})
	if err != nil {
		return err
	}

	//Purge
	_, err = amazefulDB.Collection(CollectionStr(CollectionPurge)).Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.M{"channel": 1}, Options: options.Index().SetUnique(true)},
	})
	if err != nil {
		return err
	}

	//User
	_, err = amazefulDB.Collection(CollectionStr(CollectionUser)).Indexes().CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.M{"userId": 1}, Options: options.Index().SetUnique(true)},
		{Keys: bson.M{"login": 1}},
		{Keys: bson.M{"admin": 1}},
	})
	if err != nil {
		return err
	}

	return nil

}
