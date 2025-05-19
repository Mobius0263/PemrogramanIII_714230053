package controllers

import (
	handlers "sewakendaraan/handler"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupConsumerRoutes(api fiber.Router, db *mongo.Database) {
	consumer := api.Group("/consumers")
	consumer.Get("/", handlers.GetAllConsumers(db))
	consumer.Get("/:id", handlers.GetConsumerByID(db))
	consumer.Post("/", handlers.CreateConsumer(db))
	consumer.Put("/:id", handlers.UpdateConsumer(db))
	consumer.Delete("/:id", handlers.DeleteConsumer(db))
}
