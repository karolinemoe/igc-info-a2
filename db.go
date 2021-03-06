package main

import (
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

var db *mongo.Database

func DBConnect() (bool, error) {
	connection, err := mongo.NewClient("mongodb://admin:admin1@ds247439.mlab.com:47439/igca2")

	connection.Connect(context.Background())
	if err != nil {
		return false, err
	}
	db = connection.Database("igca2")
	return true, err
}

func GetTracks() *bson.Document {
	collection := db.Collection("tracks")
	tracks := bson.NewDocument()
	//err := collection.FindOne(context.Background(), map[string]string{}).Decode(&tracks)
	test, err := collection.Find(nil, nil)
	//err := collection.Find(bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	tracks.WriteDocument(32, test)
	return tracks
}

func InsertTrack(track IGCTrack) interface{} {
	collection := db.Collection("tracks")

	res, err := collection.InsertOne(context.Background(), &track)
	if err != nil {
		log.Fatal(err)
	}
	return res.InsertedID
}

func FindTrack(ID string) (interface{}, error) {

	collection := db.Collection("tracks")

	track := bson.NewDocument()
	err := collection.FindOne(context.Background(), map[string]string{"id": ID}).Decode(&track)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	track.WriteDocument(32, &track)
	return track, nil
}
