package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func MongoConnect(uri, database string) error {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	db = client.Database(database)
	return nil
}

func GetDatabase() *mongo.Database {
	return db
}

func Close() error {
	return db.Client().Disconnect(context.Background())
}
