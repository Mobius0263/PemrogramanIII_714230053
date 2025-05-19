package controllers

import (
	"sewakendaraan/handler"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRentRoutes(api fiber.Router, db *mongo.Database) {
	Rent := api.Group("/Rents")
	Rent.Get("/", handlers.GetAllRents(db))
	Rent.Get("/:id", handlers.GetRentByID(db))
	Rent.Post("/", handlers.CreateRent(db))
	Rent.Put("/:id", handlers.UpdateRent(db))
	Rent.Delete("/:id", handlers.DeleteRent(db))
}
