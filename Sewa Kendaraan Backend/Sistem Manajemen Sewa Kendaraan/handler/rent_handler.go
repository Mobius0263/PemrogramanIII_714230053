package handlers

import (
	"context"
	"fmt"
	"log"
	"sewakendaraan/config"
	"sewakendaraan/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllRents(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("GetAllRents called")

		collection := config.GetCollection(db, config.RentCollection)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var rents []bson.M
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			log.Printf("Error in GetAllRents (Find): %v", err)
			return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("Failed to fetch rents: %v", err)})
		}
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &rents); err != nil {
			log.Printf("Error in GetAllRents (Decode): %v", err)
			return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("Failed to decode rents: %v", err)})
		}

		// Convert ObjectIDs to strings and handle potential string IDs
		for i, rent := range rents {
			if id, ok := rent["_id"].(primitive.ObjectID); ok {
				rents[i]["_id"] = id.Hex()
				log.Printf("Converted rent ID: %s", rents[i]["_id"])
			}

			if consumerID, ok := rent["consumerId"].(primitive.ObjectID); ok {
				rents[i]["consumerId"] = consumerID.Hex()
			} else if consumerIDStr, ok := rent["consumerId"].(string); ok {
				rents[i]["consumerId"] = consumerIDStr
			}
			log.Printf("Consumer ID for rent %s: %v", rents[i]["_id"], rents[i]["consumerId"])

			if vehicleID, ok := rent["vehicleId"].(primitive.ObjectID); ok {
				rents[i]["vehicleId"] = vehicleID.Hex()
			} else if vehicleIDStr, ok := rent["vehicleId"].(string); ok {
				rents[i]["vehicleId"] = vehicleIDStr
			}
			log.Printf("Vehicle ID for rent %s: %v", rents[i]["_id"], rents[i]["vehicleId"])
		}

		log.Printf("Returning %d rents", len(rents))
		return c.JSON(rents)
	}
}

func GetRentByID(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("GetRentByID called")

		collection := config.GetCollection(db, config.RentCollection)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		idParam := c.Params("id")
		log.Printf("Received ID parameter: %s", idParam)

		if idParam == "" || idParam == "undefined" {
			log.Println("Invalid or undefined ID received")
			return c.Status(400).JSON(fiber.Map{"error": "Invalid or undefined ID"})
		}

		id, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			log.Printf("Invalid ID in GetRentByID: %v", err)
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
		}

		var rent models.Rent
		err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&rent)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				log.Printf("Rent not found in GetRentByID: %v", err)
				return c.Status(404).JSON(fiber.Map{"error": "Rent not found"})
			}
			log.Printf("Error in GetRentByID: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("Failed to fetch rent: %v", err)})
		}

		log.Printf("Successfully fetched rent with ID: %s", idParam)
		return c.JSON(rent)
	}
}

func CreateRent(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("CreateRent called")

		collection := config.GetCollection(db, config.RentCollection)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var rent models.Rent
		if err := c.BodyParser(&rent); err != nil {
			log.Printf("Error parsing request body in CreateRent: %v", err)
			return c.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("Invalid request body: %v", err)})
		}

		rent.ID = primitive.NewObjectID()
		_, err := collection.InsertOne(ctx, rent)
		if err != nil {
			log.Printf("Error in CreateRent: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("Failed to create rent: %v", err)})
		}

		return c.Status(201).JSON(rent)
	}
}

func UpdateRent(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("UpdateRent called")

		collection := config.GetCollection(db, config.RentCollection)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		idParam := c.Params("id")
		id, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			log.Printf("Invalid ID in UpdateRent: %v", err)
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}

		var rent models.Rent
		if err := c.BodyParser(&rent); err != nil {
			log.Printf("Error parsing request body in UpdateRent: %v", err)
			return c.Status(400).JSON(fiber.Map{"error": fmt.Sprintf("Invalid request body: %v", err)})
		}

		update := bson.M{
			"$set": rent,
		}
		_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, update)
		if err != nil {
			log.Printf("Error in UpdateRent: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("Failed to update rent: %v", err)})
		}

		return c.SendStatus(200)
	}
}

func DeleteRent(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Println("DeleteRent called")

		collection := config.GetCollection(db, config.RentCollection)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		idParam := c.Params("id")
		id, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			log.Printf("Invalid ID in DeleteRent: %v", err)
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
		}

		result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
		if err != nil {
			log.Printf("Error in DeleteRent: %v", err)
			return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("Failed to delete rent: %v", err)})
		}

		if result.DeletedCount == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Rent not found"})
		}

		return c.SendStatus(200)
	}
}
