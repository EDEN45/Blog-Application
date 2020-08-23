package main

import (
	"context"
	"github.com/EDEN45/Blog-Application/global"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	global.MongoDB.Collection("users").InsertOne(context.Background(), bson.M{"name": "test"})
}
