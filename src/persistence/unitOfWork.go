package persistence

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type unitOfWork struct {
	context     context.Context
	transaction pgx.Tx
}

func (unitOfWork *unitOfWork) SaveChanges() error {
	return unitOfWork.transaction.Commit(unitOfWork.context)
}

func (unitOfWork *unitOfWork) Rollback() {
	//safe to ignore error
	//Rollback will be always differed to ensure that on error transaction is rolledback
	unitOfWork.transaction.Rollback(unitOfWork.context)
}
