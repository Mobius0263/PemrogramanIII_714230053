package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DBName                    = "sewakendaraan"
	VehicleCollection         = "vehicles"
	ConsumerCollection        = "consumers"
	RentCollection            = "rents"
	MongoString        string = os.Getenv("MONGODB_URI")
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	MongoString = os.Getenv("MONGODB_URI")
	if MongoString == "" {
		log.Println("Warning: MONGODB_URI not set in .env file, using default")
		MongoString = "mongodb://localhost:27017"
	}

	dbName := os.Getenv("DATABASE_NAME")
	if dbName != "" {
		DBName = dbName
	}
}

func MongoConnect() (db *mongo.Database) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if MongoString == "" {
		log.Fatal("MONGODB_URI not set in .env file")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
		return nil
	}

	fmt.Println("MongoDB connected!")
	return client.Database(DBName)
}

func GetCollection(db *mongo.Database, collectionName string) *mongo.Collection {
	return db.Collection(collectionName)
}
