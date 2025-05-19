package controllers

import (
	"sewakendaraan/handler"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupVehicleRoutes(api fiber.Router, db *mongo.Database) {
	Vehicle := api.Group("/Vehicles")
	Vehicle.Get("/", handlers.GetAllVehicles(db))
	Vehicle.Get("/:id", handlers.GetVehicleByID(db))
	Vehicle.Post("/", handlers.CreateVehicle(db))
	Vehicle.Put("/:id", handlers.UpdateVehicle(db))
	Vehicle.Delete("/:id", handlers.DeleteVehicle(db))
}
