package cobasql

import (
	"database/sql"
	"fmt"
	c "go_first/config"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	conf := c.Config()
	DB_HOST := conf.DB_HOST
	DB_PORT := conf.DB_PORT
	DB_NAME := conf.DB_NAME
	DB_USERNAME := conf.DB_USERNAME
	DB_PASSWORD := conf.DB_PASSWORD
	tz := conf.TIMEZONE

	con := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME, url.QueryEscape(tz))
	db, err := sql.Open("mysql", con)

	if err != nil {
		return nil, err
	}

	return db, nil
}
