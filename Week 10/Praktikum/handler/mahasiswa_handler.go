package handler

import (
	"Cluster0263/model"
	"Cluster0263/repository"
	"fmt"
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
			"message": "Internal Server Error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func GetMahasiswaByNPM(c *fiber.Ctx) error {
	npmStr := c.Params("npm")

	npm, err := strconv.Atoi(npmStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid NPM format",
		})
	}

	data, err := repository.GetMahasiswaByNPM(c.Context(), npm)
	if err != nil {
		msg := fmt.Sprintf("Internal server error: %s", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": msg,
		})
	}

	if data == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Data retrieved successfully",
		"data":    data,
	})
}

func InsertMahasiswa(c *fiber.Ctx) error {
	var mhs model.Mahasiswa

	if err := c.BodyParser(&mhs); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	insertedID, err := repository.InsertMahasiswa(c.Context(), mhs)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": fmt.Sprintf("Gagal menambahkan mahasiswa: %v", err),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Mahasiswa berhasil ditambahkan",
		"id":      insertedID,
		"status":  fiber.StatusCreated,
	})
}

func UpdateMahasiswa(c *fiber.Ctx) error {
	npm := c.Params("npm")
	var mhs model.Mahasiswa

	if err := c.BodyParser(&mhs); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	npmInt, err := strconv.Atoi(npm)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid NPM format: %v", err),
		})
	}

	updatedNPM, err := repository.UpdateMahasiswa(c.Context(), npmInt, mhs)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Error Update Data Mahasiswa %s : %v", npm, err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data mahasiswa berhasil diupdate",
		"npm":     updatedNPM,
		"status":  fiber.StatusOK,
	})
}

func DeleteMahasiswa(c *fiber.Ctx) error {
	npm := c.Params("npm")
	npmInt, err := strconv.Atoi(npm)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Invalid NPM format: %v", err),
		})
	}

	deletedNPM, err := repository.DeleteMahasiswa(c.Context(), npmInt)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fmt.Sprintf("Mahasiswa dengan NPM %s tidak ditemukan: %v", npm, err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Mahasiswa berhasil dihapus",
		"npm":     deletedNPM,
		"status":  fiber.StatusOK,
	})
}
