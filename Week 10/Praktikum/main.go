package main

import (
	"Cluster0263/config"
	"Cluster0263/router"
	"fmt"
	"log"
	"os"
	"strings"

	_ "Cluster0263/docs"

	"aidanwoods.dev/go-paseto"
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

// @title TES SWAGGER PEMROGRAMAN III
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/indrariksa
// @contact.email indra@ulbi.ac.id

// @host localhost:8099
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	privateKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := privateKey.Public()

	fmt.Println("Private Key (hex):", privateKey.ExportHex())
	fmt.Println("Public Key (hex):", publicKey.ExportHex())

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
