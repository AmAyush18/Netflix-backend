package controller

import (
	"context"
	"fmt"
	"log"
	"netflix-backend/model"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://amayush0818:gobackend123@golang-backend.zachn.mongodb.net/?retryWrites=true&w=majority&appName=golang-backend"

const dbName = "netflix"

const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongoDB

func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connectio success")

	collection = client.Database(dbName).Collection(colName)

	// collection instance (more of a reference)
	fmt.Println("Collection instance is ready")
}

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}
