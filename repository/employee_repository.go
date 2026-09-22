package repository

import (
	"database/sql"

	"employee-fiber/model"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{
		DB: db,
	}
}

func (r *EmployeeRepository) GetAll() ([]model.Employee, error) {
	rows, err := r.DB.Query(
		"SELECT id, name, email, position FROM employees",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employee

	for rows.Next() {
		var employee model.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
			&employee.Position,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *EmployeeRepository) GetByID(id string) (model.Employee, error) {
	var employee model.Employee

	err := r.DB.QueryRow(
		"SELECT id, name, email, position FROM employees WHERE id = @p1",
		id,
	).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Email,
		&employee.Position,
	)

	if err == sql.ErrNoRows {
		return employee, sql.ErrNoRows
	}

	return employee, err
}

func (r *EmployeeRepository) Create(employee model.Employee) (int, error) {
	var id int

	err := r.DB.QueryRow(
		`INSERT INTO employees (name, email, position)
		 OUTPUT INSERTED.id
		 VALUES (@p1, @p2, @p3)`,
		employee.Name,
		employee.Email,
		employee.Position,
	).Scan(&id)

	return id, err
}

func (r *EmployeeRepository) Update(id string, employee model.Employee) (int64, error) {
	result, err := r.DB.Exec(
		`UPDATE employees
		 SET name = @p1, email = @p2, position = @p3
		 WHERE id = @p4`,
		employee.Name,
		employee.Email,
		employee.Position,
		id,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

func (r *EmployeeRepository) Delete(id string) (int64, error) {
	result, err := r.DB.Exec(
		"DELETE FROM employees WHERE id = @p1",
		id,
	)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}