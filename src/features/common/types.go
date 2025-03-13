package common

type IUnitOfWork interface {
	SaveChanges() error
}
