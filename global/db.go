package global

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// DB hold Database connection
var MongoDB mongo.Database

func connectToMongoDb() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDbUri))

	if err != nil {
		log.Fatal("Error connect to Mongo: ", err.Error())
	}

	MongoDB = *client.Database(mongoDbName)
}
