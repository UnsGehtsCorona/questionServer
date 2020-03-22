package database

import (
	"context"
	"fmt"
	"log"
	"sync"
	config2 "wirsindcorona/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var client *mongo.Client

func ConnectMongoDb() *mongo.Client {
	cfg := config2.GetConfig().MongoDB

	once.Do(func() {
		clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", cfg.Host))
		clientOptions.Auth = &options.Credential{
			Username: cfg.Username,
			Password: cfg.Password,
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
