package handler

import (
	"database/sql"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"employee-fiber/model"
	"employee-fiber/repository"
)

type EmployeeHandler struct {
	Repository *repository.EmployeeRepository
}

func NewEmployeeHandler(repo *repository.EmployeeRepository) *EmployeeHandler {
	return &EmployeeHandler{
		Repository: repo,
	}
}

func (h *EmployeeHandler) GetAll(c *fiber.Ctx) error {
	employees, err := h.Repository.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal mengambil data employee",
		})
	}

	return c.JSON(employees)
}

func (h *EmployeeHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	employee, err := h.Repository.GetByID(id)

	if err == sql.ErrNoRows {
		return c.Status(404).JSON(fiber.Map{
			"error": "Employee tidak ditemukan",
		})
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal mengambil data employee",
		})
	}

	return c.JSON(employee)
}

func (h *EmployeeHandler) Create(c *fiber.Ctx) error {
	var employee model.Employee

	err := c.BodyParser(&employee)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	if employee.Name == "" || employee.Email == "" || employee.Position == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Name, email, dan position wajib diisi",
		})
	}

	id, err := h.Repository.Create(employee)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal menambahkan employee",
		})
	}

	employee.ID = id

	return c.Status(201).JSON(employee)
}

func (h *EmployeeHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var employee model.Employee

	err := c.BodyParser(&employee)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Data tidak valid",
		})
	}

	if employee.Name == "" || employee.Email == "" || employee.Position == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Name, email, dan position wajib diisi",
		})
	}

	rowsAffected, err := h.Repository.Update(id, employee)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal mengubah employee",
		})
	}

	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "Employee tidak ditemukan",
		})
	}

	employee.ID, _ = strconv.Atoi(id)

	return c.JSON(employee)
}

func (h *EmployeeHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	rowsAffected, err := h.Repository.Delete(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Gagal menghapus employee",
		})
	}

	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "Employee tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Employee berhasil dihapus",
	})
}
