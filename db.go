package main

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
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

