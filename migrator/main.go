package main

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"github.com/rmmbdev/crypto-price-tracker/utils"
	"log"
	"strconv"
)

var (
	dbHost       string
	dbPortNumber int
	dbName       string
	dbUsername   string
	dbPassword   string
)

func init() {
	val, err := utils.GetEnv("DB_HOST")
	if err != nil {
		log.Fatal(err)
	}
	dbHost = val

	val, err = utils.GetEnv("DB_PORT")
	if err != nil {
		log.Fatal(err)
	}
	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
	dbPortNumber = valInt

	val, err = utils.GetEnv("DB_NAME")
	if err != nil {
		log.Fatal(err)
	}
	dbName = val

	val, err = utils.GetEnv("DB_USERNAME")
	if err != nil {
		log.Fatal(err)
	}
	dbUsername = val

	val, err = utils.GetEnv("DB_PASSWORD")
	if err != nil {
		log.Fatal(err)
	}
	dbPassword = val
}

func createTable(db *sql.DB) error {
	createTableQuery := `
        CREATE TABLE IF NOT EXISTS price (
            currency VARCHAR(255) PRIMARY KEY,
            modified_at TIMESTAMP WITH TIME ZONE,            
            price FLOAT
        )
    `

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return errors.New("problem with creating table")
	}

	return nil
}

func main() {
	db, err := utils.ConnectToPostgres(dbHost, dbPortNumber, dbUsername, dbPassword, dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = createTable(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Table created or already existed!")

}
