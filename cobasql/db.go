package cobasql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func connectDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	con := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := sql.Open("mysql", con)

	if err != nil {
		return nil, err
	}

	return db, nil
}
