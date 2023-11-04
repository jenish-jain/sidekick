package mongo

import (
	"context"
	"fmt"
	"github.com/jenish-jain/sidekick/internal/helpers"
	"github.com/jenish-jain/sidekick/pkg/quantity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type repository interface {
	GetCollectionNames() []string
	GetDBStats() DbStats
	GetServerStatus() ServerStatus
	GetCollectionStats(collectionName string) CollectionStats
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

// GetServerStatus fetches mongo db stats for connected DB
// storage values scaled to MegaBytes
func (m *repositoryImpl) GetServerStatus() ServerStatus {
	var serverStatus ServerStatus
	command := bson.D{{"serverStatus", 1}}
	bsonBytes, _ := bson.Marshal(m.runMongoCommand(command))
	bson.Unmarshal(bsonBytes, &serverStatus)
	return serverStatus
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
func (m *repositoryImpl) GetCollectionStats(collectionName string) CollectionStats {
	type mongoCollStats struct {
		Name                           string             `bson:"ns"`
		UncompressedStorageSizeInBytes float64            `bson:"size"`
		DocumentCount                  int64              `bson:"count"`
		AvgObjSizeInBytes              float64            `bson:"avgObjSize"`
		CompressedStorageSizeInBytes   float64            `bson:"storageSize"`
		IndexCount                     int32              `bson:"indexes"`
		TotalIndexSizeInBytes          float64            `bson:"totalIndexSize"`
		IndexSizesInBytes              map[string]float64 `bson:"indexSizes"`
	}

	var cs mongoCollStats
	command := bson.D{{"collStats", collectionName}}
	bsonBytes, _ := bson.Marshal(m.runMongoCommand(command))
	bson.Unmarshal(bsonBytes, &cs)

	collectionStatistics := CollectionStats{
		Name:                    cs.Name,
		UncompressedStorageSize: quantity.New(cs.UncompressedStorageSizeInBytes, quantity.Bytes),
		CompressedStorageSize:   quantity.New(cs.CompressedStorageSizeInBytes, quantity.Bytes),
		DocumentCount:           cs.DocumentCount,
		AvgObjSize:              quantity.New(cs.AvgObjSizeInBytes, quantity.Bytes),
		IndexCount:              cs.IndexCount,
		TotalIndexSize:          quantity.New(cs.TotalIndexSizeInBytes, quantity.Bytes),
	}

	firstDocTimeStamp, lastDocTimeStamp := m.getFirstAndLastAccessedTimeStamps(collectionName)
	collectionStatistics.LastAccessedDate = lastDocTimeStamp.Format("2006-01-02")
	collectionStatistics.CollectionAgeInDays = m.time.GetDaysBetweenTimestamps(firstDocTimeStamp, lastDocTimeStamp)

	indexSizes := map[string]quantity.Quantity{}
	for indexName, indexSize := range cs.IndexSizesInBytes {
		indexSizes[indexName] = quantity.New(indexSize, quantity.Bytes)
	}

	return collectionStatistics
}

func (m *repositoryImpl) getFirstAndLastAccessedTimeStamps(collectionName string) (time.Time, time.Time) {
	var firstDoc mongoDocument
	var latestDoc mongoDocument
	findOptions := options.FindOneOptions{}

	findOptions.SetSort(bson.D{{"_id", 1}})
	m.database.Collection(collectionName).FindOne(context.TODO(), bson.D{}, &findOptions).Decode(&firstDoc)

	findOptions.SetSort(bson.D{{"_id", -1}})
	m.database.Collection(collectionName).FindOne(context.TODO(), bson.D{}, &findOptions).Decode(&latestDoc)

	return firstDoc.Id.Timestamp(), latestDoc.Id.Timestamp()
}

func (m *repositoryImpl) runMongoCommand(command bson.D) bson.M {
	var result bson.M
	err := m.database.RunCommand(context.TODO(), command).Decode(&result)
	if err != nil {
		fmt.Printf("error while running command for %s , error : %+v \n", command, err)
	}
	return result
}
