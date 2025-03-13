package persistence

import (
	employees "employees-import/features/employees"
	"errors"
	"time"

	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type employeeRepository struct {
	connection  *pgxpool.Pool
	transaction pgx.Tx
	context     context.Context
}

func (repository *employeeRepository) GetAll() ([]employees.Employee, error) {

	sql := `
SELECT "id", 
	"payroll_number", 
	"forenames", 
	"surname", 
	"date_of_birth", 
	"telephone_number", 
	"mobile_number", 
	"address_line_1", 
	"address_line_2", 
	"post_code",
	"email", 
	"start_date",
	"created_at",
	"updated_at"
FROM "employees"
ORDER BY "created_at"`

	rows, queryError := repository.connection.Query(repository.context, sql)

	if queryError != nil {
		return nil, queryError
	}

	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[employee])

	if err != nil {
		return nil, err
	}

	employees := []employees.Employee{}

	for _, item := range data {
		employees = append(employees, mapToEmployee(item))
	}

	return employees, nil
}

func (repository *employeeRepository) GetById(id uuid.UUID) (bool, employees.Employee, error) {

	sql := `
SELECT "id", 
	"payroll_number", 
	"forenames", 
	"surname", 
	"date_of_birth", 
	"telephone_number", 
	"mobile_number", 
	"address_line_1", 
	"address_line_2", 
	"post_code",
	"email", 
	"start_date",
	"created_at",
	"updated_at"
FROM "employees"
WHERE "id" = $1
`

	row, _ := repository.connection.Query(repository.context, sql, id)
	employee, err := pgx.CollectExactlyOneRow(row, pgx.RowToStructByName[employee])

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return false, employees.Employee{}, nil
		}
		return false, employees.Employee{}, err
	}

	return true, mapToEmployee(employee), nil
}

func (repository *employeeRepository) GetPayrollNumberToIdMap() (map[string]uuid.UUID, error) {
	sql := `
SELECT "id", 
	"payroll_number"
FROM "employees"`

	rows, queryError := repository.transaction.Query(repository.context, sql)

	if queryError != nil {
		return nil, queryError
	}

	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[idAndPayrollNumber])

	if err != nil {
		return nil, err
	}

	mapping := make(map[string]uuid.UUID)

	for _, item := range data {
		mapping[item.PayrollNumber] = item.Id
	}

	return mapping, nil
}

func (repository *employeeRepository) IsPayrollNumberExist(payrollNumber string) (bool, error) {
	sql := `
SELECT COUNT("id")
FROM "employees"
WHERE "payroll_number" = $1`

	var count int32

	err := repository.transaction.QueryRow(repository.context, sql, payrollNumber).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repository *employeeRepository) IsPayrollNumberExistExclusive(payrollNumber string, excludeId uuid.UUID) (bool, error) {
	sql := `
SELECT COUNT("id")
FROM "employees"
WHERE "payroll_number" = $1 
	AND "id" <> $2`

	var count int32

	err := repository.transaction.QueryRow(repository.context, sql, payrollNumber, excludeId).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repository *employeeRepository) IsIdExist(id uuid.UUID) (bool, error) {
	sql := `
SELECT COUNT("id")
FROM "employees"
WHERE "id" = $1`

	var count int32

	err := repository.transaction.QueryRow(repository.context, sql, id).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repository *employeeRepository) Create(employee employees.Employee) error {
	sql := `
INSERT INTO "employees"(
	"id", 
	"payroll_number", 
	"forenames", 
	"surname", 
	"date_of_birth", 
	"telephone_number", 
	"mobile_number", 
	"address_line_1", 
	"address_line_2", 
	"post_code", 
	"email", 
	"start_date",
	"created_at",
	"updated_at")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	now := time.Now()

	_, err := repository.transaction.Exec(repository.context,
		sql,
		employee.Id,
		employee.PayrollNumber,
		employee.Forenames,
		employee.Surname,
		employee.DateOfBirth,
		employee.TelephoneNumber,
		employee.MobileNumber,
		employee.AddressLine1,
		employee.AddressLine2,
		employee.Postcode,
		employee.Email,
		employee.StartDate,
		now,
		now)

	return err
}

func (repository *employeeRepository) Update(employee employees.Employee) error {
	sql := `
UPDATE "employees"
SET "payroll_number" = $1, 
	"forenames" = $2, 
	"surname" = $3, 
	"date_of_birth" = $4, 
	"telephone_number" = $5, 
	"mobile_number" = $6, 
	"address_line_1" = $7, 
	"address_line_2" = $8, 
	"post_code" = $9, 
	"email" = $10, 
	"start_date" = $11, 
	"updated_at" = $12
WHERE "id" = $13`

	_, err := repository.transaction.Exec(repository.context,
		sql,
		employee.PayrollNumber,
		employee.Forenames,
		employee.Surname,
		employee.DateOfBirth,
		employee.TelephoneNumber,
		employee.MobileNumber,
		employee.AddressLine1,
		employee.AddressLine2,
		employee.Postcode,
		employee.Email,
		employee.StartDate,
		time.Now(),
		employee.Id)

	return err
}

func (repository *employeeRepository) Delete(id uuid.UUID) error {
	sql := `
DELETE
FROM "employees"
WHERE "id" = $1`

	_, err := repository.transaction.Exec(repository.context, sql, id)

	return err
}

type employee struct {
	Id              uuid.UUID `db:"id"`
	PayrollNumber   string    `db:"payroll_number"`
	Forenames       string    `db:"forenames"`
	Surname         string    `db:"surname"`
	DateOfBirth     time.Time `db:"date_of_birth"`
	TelephoneNumber string    `db:"telephone_number"`
	MobileNumber    string    `db:"mobile_number"`
	AddressLine1    string    `db:"address_line_1"`
	AddressLine2    string    `db:"address_line_2"`
	Postcode        string    `db:"post_code"`
	Email           string    `db:"email"`
	StartDate       time.Time `db:"start_date"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type idAndPayrollNumber struct {
	Id            uuid.UUID `db:"id"`
	PayrollNumber string    `db:"payroll_number"`
}

func mapToEmployee(employee employee) employees.Employee {
	return employees.ExistingEmployee(employee.Id,
		employee.PayrollNumber,
		employee.Forenames,
		employee.Surname,
		employee.DateOfBirth,
		employee.TelephoneNumber,
		employee.MobileNumber,
		employee.AddressLine1,
		employee.AddressLine2,
		employee.Postcode,
		employee.Email,
		employee.StartDate)
}
