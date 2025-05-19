package handler

import (
	"Cluster0263/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Homepage(c *fiber.Ctx) error {
	return c.SendString("Welcome to the homepage!")
}

func GetAllMahasiswa(c *fiber.Ctx) error {
	data, err := repository.GetAllMahasiswa(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve data", //Gagal mengambil data dari database
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK, //Status 200 (OK)
		"message": "Successfully retrieved data",
		"data":    data,
	})
}

func GetMahasiswaByNPM(c *fiber.Ctx) error {
	npmStr := c.Params("npm") //Mengambil NPM dari parameter URL

	npm, err := strconv.Atoi(npmStr) //Mengubah string NPM menjadi integer
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "NPM harus berupa angka", //Format NPM tidak valid
		})
	}

	mhs, err := repository.GetMahasiswaByNPM(c.Context(), npm)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if mhs == nil { //Jika NPM tidak ditemukan
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan", //Mahasiswa tidak ditemukan
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK, //Status 200 (OK)
		"message": "Successfully retrieved data",
		"data":    mhs,
	})
}
