package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

func init() {
	uuid.EnableRandPool()
}
func ConnectToPostgres(host string, port int, username string, password string, dbName string) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		username,
		password,
		dbName,
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GenerateUUID() (string, error) {
	res, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("problem with creating new uuid")
	}
	return res.String(), nil
}
