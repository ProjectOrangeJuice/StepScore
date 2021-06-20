package db

import (
	"time"

	"github.com/ProjectOrangeJuice/golib-logging/logging"
)

func GetBoard(code string, date time.Time) {
	stmt, err := dbConn.Prepare(`SELECT * FROM 'boards' JOIN userBoards 
	JOIN steps WHERE userBoards.boardID = ? AND boards.boardID = ? AND 
	steps.userID = userBoards.userID`)
	if err != nil {
		logging.Log("GetBoard", "sql script error", err)
	}
	defer stmt.Close()

	ttime := date.Format("02-01-2006")
	_, err = stmt.Exec(code, ttime)
	if err != nil {
		logging.Log("GetBoard", "sql error", err)
	}

}

func GetBoards(userID int) {
	stmt, err := dbConn.Prepare(`SELECT * FROM 'boards' WHERE userID=?`)
	if err != nil {
		logging.Log("GetBoards", "sql script error", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID)
	if err != nil {
		logging.Log("GetBoards", "sql error", err)
	}

}
