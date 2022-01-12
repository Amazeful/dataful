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

var dbSetupData = []struct {
	db      DBName
	col     Collection
	indexes []mongo.IndexModel
}{
	// ==============================|| AMAZEFULDB ||============================== //

	//Alerts
	{DBAmazeful, CollectionAlerts, []mongo.IndexModel{{Keys: bson.M{"channel": 1}, Options: options.Index().SetUnique(true)}}},

	//Channel
	{DBAmazeful, CollectionChannel, []mongo.IndexModel{{Keys: bson.M{"channelId": 1}, Options: options.Index().SetUnique(true)}, {Keys: bson.M{"broadcasterName": 1}}, {Keys: bson.M{"joined": 1}}}},

	//Command
	{DBAmazeful, CollectionCommand, []mongo.IndexModel{{Keys: bson.D{{"channel", 1}, {"name", 1}}, Options: options.Index().SetUnique(true)}}},

	//Filters
	{DBAmazeful, CollectionFilters, []mongo.IndexModel{{Keys: bson.M{"channel": 1}, Options: options.Index().SetUnique(true)}}},

	//Purge
	{DBAmazeful, CollectionFilters, []mongo.IndexModel{{Keys: bson.M{"channel": 1}, Options: options.Index().SetUnique(true)}}},

	//User
	{DBAmazeful, CollectionUser, []mongo.IndexModel{{Keys: bson.M{"userId": 1}, Options: options.Index().SetUnique(true)}, {Keys: bson.M{"login": 1}}, {Keys: bson.M{"admin": 1}}}},
}

//createCollectionsAndIndexes creates all collections and ensures indexes are added in.
func createCollectionsAndIndexes(ctx context.Context, c *mongo.Client) error {
	var err error
	for _, s := range dbSetupData {
		if s.indexes == nil {
			//only if no indexes
			err = c.Database(string(s.db)).CreateCollection(ctx, string(s.col))
			if err != nil {
				return err
			}
		} else {
			//Calling create indexes will automatically create collection if DNE
			_, err = c.Database(string(s.db)).Collection(string(s.col)).Indexes().CreateMany(ctx, s.indexes)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
