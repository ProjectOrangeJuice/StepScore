package db

import (
	"database/sql"
	"io/ioutil"

	"github.com/ProjectOrangeJuice/golib-logging/logging"
	_ "github.com/mattn/go-sqlite3"
)

var dbConn *sql.DB

func init() {
	var err error
	dbConn, err = sql.Open("sqlite3", "./db/script/db.sqlite")
	if err != nil {
		logging.Log("db-init", "Failed to open sqlite", err)
	}

	// Read setup script
	dat, err := ioutil.ReadFile("db/script/setup.sql")
	if err != nil {
		logging.Log("db-init", "Failed to open sql", err)
	}

	_, err = dbConn.Exec(string(dat))
	if err != nil {
		logging.Log("db-init", "Failed to read sql and execute", err)
	}

}
