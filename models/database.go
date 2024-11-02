package models

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// Инициализация  подключение к базе Данных PostgreSQL на базе библиотеки github.com/lib/pq
func DbInitialization() (*sql.DB, error) {

	// dbAddr, _ := os.LookupEnv("MB_CONFIG_DB_HOST")
	// dbPort, _ := os.LookupEnv("MB_CONFIG_DB_PORT")
	// dbUser, _ := os.LookupEnv("MB_CONFIG_DB_USER")
	// dbPassword, _ := os.LookupEnv("MB_CONFIG_DB_PASS")
	// dbName, _ := os.LookupEnv("MB_CONFIG_DB_NAME")
	// dbSchema, _ := os.LookupEnv("MB_CONFIG_DB_SCHEMA")
	// dbSslmode, _ := os.LookupEnv("MB_CONFIG_DB_SSLMODE")

	dbAddr := "localhost"
	dbPort := "5433"
	dbUser := "postgres"
	dbPassword := "postgres"
	dbName := "mbclients"
	dbSchema := "mbclients"
	dbSslmode := "disable"

	//	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s schema=%s sslmode=%s", dbAddr, dbPort, dbUser, dbPassword, dbName, dbSchema, dbSslmode)

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s schema=%s sslmode=%s", dbAddr, dbPort, dbUser, dbPassword, dbName, dbSchema, dbSslmode)) //connectionString)
	db.SetMaxOpenConns(100)

	if err != nil {
		return nil, err
	}

	return db, nil
}

var database *sql.DB

func init() {
	db, err := DbInitialization()
	err = db.Ping()
	if err != nil {
		fmt.Println(time.Now(), "DB:", err)
	} else {
		fmt.Println(time.Now(), ": DB connected")
	}
	database = db
}
