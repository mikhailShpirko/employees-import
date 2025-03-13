package persistence

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateEmployeeRepository(connectionString string, context context.Context) (*employeeRepository, *unitOfWork, error) {

	pool, err := pgxpool.New(context, connectionString)

	if err != nil {
		return nil, nil, err
	}

	transaction, err := pool.Begin(context)

	if err != nil {
		return nil, nil, err
	}

	repository := employeeRepository{}
	unitOfWork := unitOfWork{}

	repository.connection = pool
	repository.transaction = transaction
	repository.context = context

	unitOfWork.transaction = transaction
	unitOfWork.context = context

	return &repository, &unitOfWork, nil
}
