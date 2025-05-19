package routes

import (
	handlers "sewakendaraan/handler"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func RentRoutes(api fiber.Router, db *mongo.Database) {
	// Rent routes
	Rents := api.Group("/Rents")
	Rents.Get("/", handlers.GetAllRents(db))
	Rents.Get("/:id", handlers.GetRentByID(db))
	Rents.Post("/", handlers.CreateRent(db))
	Rents.Put("/:id", handlers.UpdateRent(db))
	Rents.Delete("/:id", handlers.DeleteRent(db))
}
