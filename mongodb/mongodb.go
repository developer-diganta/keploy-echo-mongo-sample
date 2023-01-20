package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Person struct
type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var client *mongo.Client

// ConnectToMongoDB connects to the MongoDB server
func ConnectToMongoDB() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
}

// DisconnectFromMongoDB disconnects from the MongoDB server
func DisconnectFromMongoDB() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from MongoDB.")
}

// InsertPerson inserts a person into the MongoDB collection
func InsertPerson(person Person) error {
	collection := client.Database("sample-app").Collection("people")
	_, err := collection.InsertOne(context.TODO(), person)

	if err != nil {
		return err
	}
	return nil
}