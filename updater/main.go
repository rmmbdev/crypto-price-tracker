package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"github.com/rmmbdev/crypto-price-tracker/utils"
	"log"
	"net/url"
	"strconv"
	"time"
)

var (
	dbHost       string
	dbPortNumber int
	dbName       string
	dbUsername   string
	dbPassword   string
	sourceAddr   string
	sourcePath   string
	sourceQuery  string
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

	val, err = utils.GetEnv("SOURCE_ADDRESS")
	if err != nil {
		log.Fatal(err)
	}
	sourceAddr = val

	val, err = utils.GetEnv("SOURCE_PATH")
	if err != nil {
		log.Fatal(err)
	}
	sourcePath = val

	val, err = utils.GetEnv("SOURCE_QUERY")
	if err != nil {
		log.Fatal(err)
	}
	sourceQuery = val

}

func updateCurrencies(db *sql.DB, data map[string]string, now time.Time) error {
	for k, v := range data {
		query := `INSERT INTO price (currency, modified_at,price)
				    VALUES ($1, $2, $3)
					ON CONFLICT (currency)
					DO
					  UPDATE SET price = EXCLUDED.price`

		_, err := db.Exec(query, k, now, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	db, err := utils.ConnectToPostgres(dbHost, dbPortNumber, dbUsername, dbPassword, dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.SetFlags(0)

	u := url.URL{Scheme: "wss", Host: sourceAddr, Path: sourcePath, RawQuery: sourceQuery}
	log.Printf("connecting to %s", u.String())

	// Dial the WebSocket connection
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// Read messages from the WebSocket connection
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
			return
		}

		var data map[string]string
		err = json.Unmarshal(message, &data)
		if err != nil {
			log.Fatal("decode:", err)
			return
		}
		//log.Println("data", data)
		now := time.Now().UTC()
		err = updateCurrencies(db, data, now)
		if err != nil {
			log.Fatal("insert:", err)
		}
		log.Println("Data inserted to database at [", now.String(), "]")
		log.Println("**************************************************")
	}
}
