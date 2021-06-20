package db

import (
	"github.com/ProjectOrangeJuice/golib-logging/logging"
)

func InsertSteps(userID, steps int, date string) {
	stmt, err := dbConn.Prepare(`INSERT OR REPLACE INTO 'steps' 
	('userID','date','steps') VALUES (?, ?, ?);`)
	if err != nil {
		logging.Log("insertSteps", "sql script error", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, date, steps)
	if err != nil {
		logging.Log("insertSteps", "sql error", err)
	}

}

func InsertBoard(code, description string) {
	stmt, err := dbConn.Prepare(`INSERT INTO 'boards' 
	('boardID','description') VALUES (?, ?);`)
	if err != nil {
		logging.Log("InsertBoard", "sql script error", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(code, description)
	if err != nil {
		logging.Log("InsertBoard", "sql error", err)
	}

}

func JoinBoard(userID int, code string) {
	stmt, err := dbConn.Prepare(`INSERT INTO 'userBoards' 
	('userID','boardID') VALUES (?, ?);`)
	if err != nil {
		logging.Log("JoinBoard", "sql script error", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, code)
	if err != nil {
		logging.Log("JoinBoard", "sql error", err)
	}

}
