package handlers

import (
	"context"
	"sewakendaraan/config"
	"sewakendaraan/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllVehicles(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := config.GetCollection(db, config.VehicleCollection)
		cursor, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		var Vehicles []models.Vehicle
		if err = cursor.All(context.Background(), &Vehicles); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(Vehicles)
	}
}

func GetVehicleByID(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := config.GetCollection(db, config.VehicleCollection)
		idParam := c.Params("id")
		id, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}
		var Vehicle models.Vehicle
		err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&Vehicle)
		if err != nil {
			return c.Status(404).SendString("Vehicle not found")
		}
		return c.JSON(Vehicle)
	}
}

func CreateVehicle(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := config.GetCollection(db, config.VehicleCollection)
		var Vehicle models.Vehicle
		if err := c.BodyParser(&Vehicle); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		Vehicle.ID = primitive.NewObjectID()
		_, err := collection.InsertOne(context.Background(), Vehicle)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(Vehicle)
	}
}

func UpdateVehicle(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := config.GetCollection(db, config.VehicleCollection)
		idParam := c.Params("id")
		id, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}
		var Vehicle models.Vehicle
		if err := c.BodyParser(&Vehicle); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		update := bson.M{
			"$set": Vehicle,
		}
		_, err = collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.SendStatus(200)
	}
}

func DeleteVehicle(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := config.GetCollection(db, config.VehicleCollection)
		idParam := c.Params("id")
		id, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}
		_, err = collection.DeleteOne(context.Background(), bson.M{"_id": id})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.SendStatus(200)
	}
}
