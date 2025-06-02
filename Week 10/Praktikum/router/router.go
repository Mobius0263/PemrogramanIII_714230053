package router

import (
	"Cluster0263/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", handler.Homepage)
	api.Get("/mahasiswa", handler.GetAllMahasiswa)
	api.Get("/mahasiswa/:npm", handler.GetMahasiswaByNPM)
	api.Post("/mahasiswa", handler.InsertMahasiswa)
	api.Put("/mahasiswa/:npm", handler.UpdateMahasiswa)
	api.Delete("/mahasiswa/:npm", handler.DeleteMahasiswa)
}
