package route

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"wirsindcorona/database"
)

var questionsColl *mongo.Collection
var answerColl *mongo.Collection

func InitCollections() {
	// Init Questions Collection
	questionsColl = database.GetDatabase().Collection("questionsColl")
	if _, err := questionsColl.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{"quid", 1}},
			Options: options.Index().SetUnique(true),
		},
	); err != nil {
		log.Fatal(err)
	}

	// Init Answer Collection
	answerColl = database.GetDatabase().Collection("answerColl")
}
