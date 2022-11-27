package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewMongoClient(uri string) *mongo.Client {
	fmt.Println("Establishing connection to mongodb")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		fmt.Errorf("error creating mongo client %s", err.Error())
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	connectErr := client.Connect(ctx)
	if connectErr != nil {
		fmt.Errorf("erorr pinging to connected db")
		panic(err)
	}
	return client
}
