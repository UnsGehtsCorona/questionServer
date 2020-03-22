package database

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var client *mongo.Client

func ConnectMongoDb() *mongo.Client {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
		clientOptions.Auth = &options.Credential{
			Username: "root",
			Password: "wirsindcorona",
		}

		c, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		err = c.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}

		client = c
	})

	return client
}

func GetDatabase() *mongo.Database {
	return client.Database("wirsindcorona")
}

func Disconnect() {
	if client != nil {
		err := client.Disconnect(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
	}
}
