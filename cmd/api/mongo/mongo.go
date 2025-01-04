package mongo

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetMongoClient() (*mongo.Client, func()) {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://db:27017"))
	if err != nil {
		log.Fatalln("Failed to connect to mongo")
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalln("Could not connect to MongoDB: ", err)
	}

	return client, func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatalln("Failed to disconnect from mongo")
			panic(err)
		}
	}
}
