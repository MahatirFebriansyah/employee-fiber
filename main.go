// package main

// import (
// 	"fmt"
// 	"github.com/gofiber/fiber/v2"
// 	"strconv"
// )

// func main() {
// 	connectDB()

// 	app := fiber.New()

// 	app.Get("/employees", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		rows, err := db.Query("SELECT id, name, email, position FROM employees WHERE id = @p1", id)
// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{
// 				"error": "Gagal mengambil data employee",
// 			})
// 		}
// 		defer rows.Close()

// 		var employees []Employee

// 		for rows.Next() {
// 			var employee Employee

// 			err := rows.Scan(
// 				&employee.ID,
// 				&employee.Name,
// 				&employee.Email,
// 				&employee.Position,
// 			)

// 			if err != nil {
// 				return c.Status(500).JSON(fiber.Map{
// 					"error": "Gagal membaca data employee",
// 				})
// 			}

// 			employees = append(employees, employee)
// 		}

// 		return c.JSON(employees)
// 	})

// 	app.Get("/employees/search", func(c *fiber.Ctx) error {
// 		name := c.Query("name")

// 		var employee Employee

// 		err := db.QueryRow(
// 			"SELECT id, name, email, position FROM employees WHERE name = @p1",
// 			name,
// 		).Scan(
// 			&employee.ID,
// 			&employee.Name,
// 			&employee.Email,
// 			&employee.Position,
// 		)

// 		if err != nil {
// 			return c.Status(404).JSON(fiber.Map{
// 				"error": "Employee tidak ditemukan",
// 			})
// 		}

// 		return c.JSON(employee)
// 	})

// 	app.Get("/employees/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")

// 		var employee Employee

// 		err := db.QueryRow(
// 			"SELECT id, name, email, position FROM employees WHERE id = @p1",
// 			id,
// 		).Scan(
// 			&employee.ID,
// 			&employee.Name,
// 			&employee.Email,
// 			&employee.Position,
// 		)

// 		if err != nil {
// 			return c.Status(404).JSON(fiber.Map{
// 				"error": "Employee tidak ditemukan",
// 			})
// 		}

// 		return c.JSON(employee)
// 	})

// 	app.Post("/employees", func(c *fiber.Ctx) error {
// 		var employee Employee

// 		err := c.BodyParser(&employee)
// 		if err != nil {
// 			return c.Status(400).JSON(fiber.Map{
// 				"error": "Data tidak valid",
// 			})
// 		}

// 		var id int

// 		err = db.QueryRow(
// 			`INSERT INTO employees (name, email, position)
// 		 OUTPUT INSERTED.id
// 		 VALUES (@p1, @p2, @p3)`,
// 			employee.Name,
// 			employee.Email,
// 			employee.Position,
// 		).Scan(&id)

// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{
// 				"error": "Gagal menambahkan employee",
// 			})
// 		}

// 		employee.ID = id

// 		return c.Status(201).JSON(employee)
// 	})

// 	app.Put("/employees/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")
// 		var employee Employee

// 		err := c.BodyParser(&employee)
// 		if err != nil {
// 			return c.Status(400).JSON(fiber.Map{
// 				"error": "Data tidak valid",
// 			})
// 		}

// 		result, err := db.Exec(
// 			"UPDATE employees SET name = @p1, email = @p2, position = @p3 WHERE id = @p4",
// 			employee.Name,
// 			employee.Email,
// 			employee.Position,
// 			id,
// 		)

// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{
// 				"error": "Gagal memperbarui employee",
// 			})
// 		}

// 		rowsAffected, err := result.RowsAffected()
// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{
// 				"error": "Gagal memperbarui employee",
// 			})
// 		}

// 		if rowsAffected == 0 {
// 			return c.Status(404).JSON(fiber.Map{
// 				"error": "Employee tidak ditemukan",
// 			})
// 		}

// 		employee.ID, _ = strconv.Atoi(id)

// 		return c.JSON(employee)
// 	})

// 	app.Delete("/employees/:id", func(c *fiber.Ctx) error {
// 		id := c.Params("id")

// 		result, err := db.Exec(
// 			"DELETE FROM employees WHERE id = @p1",
// 			id,
// 		)

// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{
// 				"error": "Gagal menghapus employee",
// 			})
// 		}

// 		rowsAffected, err := result.RowsAffected()
// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{
// 				"error": "Gagal mengecek data",
// 			})
// 		}

// 		if rowsAffected == 0 {
// 			return c.Status(404).JSON(fiber.Map{
// 				"error": "Employee tidak ditemukan",
// 			})
// 		}

// 		return c.JSON(fiber.Map{
// 			"message": "Employee berhasil dihapus",
// 		})
// 	})

// 	fmt.Println("Server is running on http://localhost:3000")
// 	app.Listen(":3000")
// }

package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"employee-fiber/handler"
	"employee-fiber/repository"
)

func main() {
	connectDB()

	app := fiber.New()

	employeeRepository := repository.NewEmployeeRepository(db)
	employeeHandler := handler.NewEmployeeHandler(employeeRepository)

	app.Get("/employees", employeeHandler.GetAll)
	app.Get("/employees/:id", employeeHandler.GetByID)
	app.Post("/employees", employeeHandler.Create)
	app.Put("/employees/:id", employeeHandler.Update)
	app.Delete("/employees/:id", employeeHandler.Delete)

	fmt.Println("Employee API is running on http://localhost:3000")

	app.Listen(":3000")
}