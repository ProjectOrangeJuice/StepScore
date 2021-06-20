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
	dbConn, err = sql.Open("sqlite3", "./db/src/db.sqlite")
	if err != nil {
		logging.Log("db-init", "Failed to open sqlite", err)
	}

	// Read setup script
	dat, err := ioutil.ReadFile("db/src/setup.sql")
	if err != nil {
		logging.Log("db-init", "Failed to open sql", err)
	}

	_, err = dbConn.Exec(string(dat))
	if err != nil {
		logging.Log("db-init", "Failed to read sql and execute", err)
	}

}

func GetUserID(username string) int {
	stmt, err := dbConn.Prepare(`SELECT userID FROM 'users' WHERE username=?`)
	if err != nil {
		logging.Log("getUserID", "sql script error", err)
		return -1
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		logging.Log("getUserID", "sql error", err)
		return -1
	}

	for rows.Next() {
		var code int

		rows.Scan(&code)
		return code
	}

	return -1

}

func AddUser(username string) {
	stmt, err := dbConn.Prepare(`INSERT INTO 'users' 
	('username') VALUES (?);`)
	if err != nil {
		logging.Log("addUser", "sql script error", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(username)
	if err != nil {
		logging.Log("addUser", "sql error", err)
		return
	}

}
