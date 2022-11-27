package mongo

import (
	"context"
	"fmt"
	"github.com/jenish-jain/sidekick/internal/helpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repository interface {
	GetCollectionNames() []string
	GetDBStats() DbStats
	GetCollectionStats(collectionName string) CollectionStats
	GetCollectionAgeInDays(collectionName string) int64
}

type repositoryImpl struct {
	database *mongo.Database
	time     helpers.Time
}

func InitMongoRepository(database *mongo.Database, time helpers.Time) repository {
	return &repositoryImpl{
		database: database,
		time:     time,
	}
}

func (m *repositoryImpl) GetCollectionNames() []string {
	filter := bson.D{{}}
	listOfColl, error := m.database.ListCollectionNames(context.TODO(), filter)
	if error != nil {
		fmt.Errorf("error listing collection names for database, error : %s", error.Error())
		panic(error)
	}
	return listOfColl
}

// GetDBStats fetches mongo db stats for connected DB
// storage values scaled to MegaBytes
func (m *repositoryImpl) GetDBStats() DbStats {
	var dbStats DbStats
	command := bson.D{{"dbStats", 1}}
	bsonBytes, _ := bson.Marshal(m.runMongoCommand(command))
	bson.Unmarshal(bsonBytes, &dbStats)
	return dbStats
}

// GetCollectionStats fetches mongo db collection stats for the specified
// collection name storage values scaled to MegaBytes
func (m *repositoryImpl) GetCollectionStats(collectionName string) CollectionStats {
	var collStats CollectionStats
	command := bson.D{{"collStats", collectionName}}
	bsonBytes, _ := bson.Marshal(m.runMongoCommand(command))
	bson.Unmarshal(bsonBytes, &collStats)
	return collStats
}

func (m *repositoryImpl) GetCollectionAgeInDays(collectionName string) int64 {
	var firstDoc mongoDocument
	var latestDoc mongoDocument
	findOptions := options.FindOneOptions{}

	findOptions.SetSort(bson.D{{"_id", 1}})
	m.database.Collection(collectionName).FindOne(context.TODO(), bson.D{}, &findOptions).Decode(&firstDoc)

	findOptions.SetSort(bson.D{{"_id", -1}})
	m.database.Collection(collectionName).FindOne(context.TODO(), bson.D{}, &findOptions).Decode(&latestDoc)

	return m.time.GetDaysBetweenTimestamps(firstDoc.Id.Timestamp(), latestDoc.Id.Timestamp())
}

func (m *repositoryImpl) runMongoCommand(command bson.D) bson.M {
	var result bson.M
	err := m.database.RunCommand(context.TODO(), command).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result
}
