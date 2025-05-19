package routes

import (
	handlers "sewakendaraan/handler"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConsumerRoutes(api fiber.Router, db *mongo.Database) {
	// Consumer routes
	consumers := api.Group("/consumers")
	consumers.Get("/", handlers.GetAllConsumers(db))
	consumers.Get("/:id", handlers.GetConsumerByID(db))
	consumers.Post("/", handlers.CreateConsumer(db))
	consumers.Put("/:id", handlers.UpdateConsumer(db))
	consumers.Delete("/:id", handlers.DeleteConsumer(db))
}
