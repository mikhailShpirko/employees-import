package configuration

import (
	"os"
)

func GetDatabaseConnectionString() string {
	return os.Getenv("DATABASE_CONNECTION_STRING")
}
