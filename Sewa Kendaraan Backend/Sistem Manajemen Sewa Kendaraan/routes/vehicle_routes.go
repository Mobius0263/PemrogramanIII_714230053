package routes

import (
	"sewakendaraan/handler"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func VehicleRoutes(api fiber.Router, db *mongo.Database) {
	// Vehicle routes
	Vehicles := api.Group("/Vehicles")
	Vehicles.Get("/", handlers.GetAllVehicles(db))
	Vehicles.Get("/:id", handlers.GetVehicleByID(db))
	Vehicles.Post("/", handlers.CreateVehicle(db))
	Vehicles.Put("/:id", handlers.UpdateVehicle(db))
	Vehicles.Delete("/:id", handlers.DeleteVehicle(db))
}
