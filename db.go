package main

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

var db *mongo.Database
var collection = db.Collection("tracks")


func DBConnect() (bool, error) {
	connection, err := mongo.NewClient("mongodb://admin:admin1@ds247439.mlab.com:47439/igca2")

	connection.Connect(context.Background())
	if err != nil {
		return false, err
	}
	db = connection.Database("igca2")
	return true, err
}

func InsertTrack(track IGCTrack) interface{} {
	res, err := collection.InsertOne(context.Background(), &track)
	if err != nil {
		log.Fatal(err)
	}
	return res.InsertedID
}