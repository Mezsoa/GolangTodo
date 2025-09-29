package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
	"log"
	"os"
	"time"
)

var (
	Client         *mongo.Client
	UsersColl      *mongo.Collection
	databaseName   = getEnv("MONGO_DB", "todolist")
	usersCollName  = getEnv("MONGO_USERS_COLL", "users")
	defaultConnURI = "mongodb://localhost:27017"
)

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func ConnectMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := getEnv("MONGO_URI", defaultConnURI)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	// Ping
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	Client = client
	UsersColl = client.Database(databaseName).Collection(usersCollName)
	log.Printf("Connected to MongoDB, db=%s coll=%s\", databaseName, usersCollName")
	return nil
}

func CloseMongo() {
	if Client != nil {
		_ = Client.Disconnect(context.Background())
	}
}
