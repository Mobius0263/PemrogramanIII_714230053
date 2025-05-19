package main

import (
	"Cluster0263/config"
	"Cluster0263/router"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env files:", err)
	}
}

func main() {
	app := fiber.New()

	//Logging Request di terminal
	app.Use(logger.New())

	//Basic CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(config.GetAllowedOrigins(), ","),
		AllowCredentials: true,
		AllowMethods:     "GET,POST,PUT,DELETE",
	}))

	//Route Mahasiswa
	router.SetupRoutes(app)

	//Handler 404
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Endpoint not found",
		})
	})

	//Baca PORT yang ada di .env
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" //Default port
	}

	//Untuk log cek koneksi port
	fmt.Printf("Server running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
		//Koneksi terputus
	}
}
