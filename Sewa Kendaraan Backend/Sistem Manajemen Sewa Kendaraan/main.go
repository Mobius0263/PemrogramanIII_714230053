package main

import (
	"context"
	"fmt"
	"log"
	"sewakendaraan/config"
	"sewakendaraan/routes"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// Connect to MongoDB
	db := config.MongoConnect()
	if db == nil {
		log.Fatal("Failed to connect to database")
	}

	// Initialize collections
	err := initializeCollections(db)
	if err != nil {
		log.Fatalf("Failed to initialize collections: %v", err)
	}

	app := fiber.New()

	// Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))
	app.Use(logger.New())

	// Routes
	api := app.Group("/api")
	setupRoutes(api, db)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

func setupRoutes(api fiber.Router, db *mongo.Database) {
	routes.VehicleRoutes(api, db)
	routes.ConsumerRoutes(api, db)
	routes.RentRoutes(api, db)
}

func initializeCollections(db *mongo.Database) error {
	// List of collections to create
	collections := []string{
		config.VehicleCollection,
		config.ConsumerCollection,
		config.RentCollection,
	}

	for _, collName := range collections {
		err := db.CreateCollection(context.Background(), collName)
		if err != nil {
			// If the collection already exists, it's not an error
			if strings.Contains(err.Error(), "already exists") {
				log.Printf("Collection %s already exists\n", collName)
			} else {
				return fmt.Errorf("failed to create collection %s: %v", collName, err)
			}
		} else {
			log.Printf("Collection %s created successfully\n", collName)
		}
	}

	// Optionally, create indexes here if needed
	// For example:
	// _, err := db.Collection(config.VehicleCollection).Indexes().CreateOne(
	//     context.Background(),
	//     mongo.IndexModel{
	//         Keys: bson.D{{Key: "brand", Value: 1}},
	//     },
	// )
	// if err != nil {
	//     return fmt.Errorf("failed to create index on vehicles collection: %v", err)
	// }

	return nil
}
