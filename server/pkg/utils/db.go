package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func ConnectDB(databaseHost, databasePort, databaseUser, databasePassword, databaseName string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		databaseHost,
		databasePort,
		databaseUser,
		databasePassword,
		databaseName,
	)

	dbInstance, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to Database")
	}
	err = dbInstance.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to Database")
	}
	return dbInstance, nil
}
