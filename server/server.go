package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"wecal/db"
	"wecal/web"
)

var connectionString = os.Getenv("DB_CONNECTION")

func main() {
	client, err := mongo.Connect(context.TODO(), clientOptions())
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())
	mongoDB := db.NewMongo(client)
	// CORS is enabled only in prod profile
	cors := os.Getenv("profile") == "prod"
	app := web.NewApp(mongoDB, cors)
	log.Println("connectionString: ", connectionString)
	err = app.Serve()
	log.Println("Error", err)
}

func clientOptions() *options.ClientOptions {
	// if os.Getenv("profile") != "prod" {
	// 	connectionString = "mongodb://localhost:27017"
	// }
	return options.Client().ApplyURI(connectionString)
}
