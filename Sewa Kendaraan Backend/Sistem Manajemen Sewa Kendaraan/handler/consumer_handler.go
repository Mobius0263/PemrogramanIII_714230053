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

func GetAllConsumers(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := config.GetCollection(db, config.ConsumerCollection)
		cursor, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		var consumers []models.Consumer
		if err = cursor.All(context.Background(), &consumers); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(consumers)
	}
}

func GetConsumerByID(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := config.GetCollection(db, config.ConsumerCollection)
		idParam := c.Params("id")
		id, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}
		var consumer models.Consumer
		err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&consumer)
		if err != nil {
			return c.Status(404).SendString("Consumer not found")
		}
		return c.JSON(consumer)
	}
}

func CreateConsumer(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := config.GetCollection(db, config.ConsumerCollection)
		var consumer models.Consumer
		if err := c.BodyParser(&consumer); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		consumer.ID = primitive.NewObjectID()
		_, err := collection.InsertOne(context.Background(), consumer)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(consumer)
	}
}

func UpdateConsumer(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := config.GetCollection(db, config.ConsumerCollection)
		idParam := c.Params("id")
		id, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			return c.Status(400).SendString("Invalid ID")
		}
		var consumer models.Consumer
		if err := c.BodyParser(&consumer); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		update := bson.M{
			"$set": consumer,
		}
		_, err = collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.SendStatus(200)
	}
}

func DeleteConsumer(db *mongo.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		collection := config.GetCollection(db, config.ConsumerCollection)
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
