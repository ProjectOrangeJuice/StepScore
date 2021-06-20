package db

import (
	"github.com/ProajgeOrangeJuice/StepCounter/data"
	"github.com/ProjectOrangeJuice/golib-logging/logging"
)

func GetBoard(code string, date string) []data.BoardResult {
	stmt, err := dbConn.Prepare(`SELECT userBoards.userID, steps.steps FROM 'boards' JOIN userBoards 
	JOIN steps WHERE userBoards.boardID = ? AND boards.boardID = ? AND 
	steps.userID = userBoards.userID and steps.date = ?`)
	if err != nil {
		logging.Log("GetBoard", "sql script error", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(code, code, date)
	if err != nil {
		logging.Log("GetBoard", "sql error", err)
	}

	board := make([]data.BoardResult, 0)

	for rows.Next() {
		var user, steps int

		rows.Scan(&user, &steps)
		board = append(board, data.BoardResult{UserID: user, Steps: steps})
	}
	return board
}

func GetBoards(userID int) []string {
	stmt, err := dbConn.Prepare(`SELECT * FROM 'userBoards' WHERE userID=?`)
	if err != nil {
		logging.Log("GetBoards", "sql script error", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		logging.Log("GetBoards", "sql error", err)
	}

	boards := make([]string, 0)

	for rows.Next() {
		var code string

		rows.Scan(&code)
		boards = append(boards, code)
	}

	return boards

}
