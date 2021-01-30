package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

var Db *mongo.Database

func CreateClient() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_CONNECT")))

	if err != nil {
		log.Println("[FATAL] could not create client for database")
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	err = client.Connect(ctx)

	defer cancel()

	if err != nil {
		log.Fatal("[FATAL] Couldn't connect to database")
	}

	Db = client.Database("cluster0")

	return
}
