package main

import (
	"encoding/json"
	"net/http"

	"github.com/keploy/go-sdk/integrations/kecho/v4"

	"sample-app/mongodb"

	"github.com/keploy/go-sdk/integrations/kmongo"
	"github.com/keploy/go-sdk/keploy"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var col *kmongo.Collection
func main() {
	// Connect to MongoDB
	mongodb.ConnectToMongoDB()
	defer mongodb.DisconnectFromMongoDB()
	dbName, collection := "sample-app", "people"
	// create a new client with mongodb

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	
	if err != nil {
		panic(err)
	}

	db := client.Database(dbName)
col = kmongo.NewCollection(db.Collection(collection))
	// Create a new Echo server
	e := echo.New()
		port:="8080"
	k := keploy.New(keploy.Config{
	App: keploy.AppConfig{
		Name: "example",
		Port: port,
	},
	Server: keploy.ServerConfig{
		URL: "http://localhost:6789/api",
	},
	})

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// kecho.EchoV4(k, e)
e.Use(kecho.EchoMiddlewareV4(k))
	// Define routes
	e.POST("/person", createPerson)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}

func createPerson(c echo.Context) error {
	// Get the name and email from the request body
	var person mongodb.Person
	defer c.Request().Body.Close()
	err := json.NewDecoder(c.Request().Body).Decode(&person)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Insert the person into the MongoDB collection
	if err := mongodb.InsertPerson(person); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Person added successfully!"})
}