package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/ProajgeOrangeJuice/StepCounter/db"
	"github.com/ProjectOrangeJuice/golib-logging/logging"
	"github.com/gorilla/mux"
)

func joinBoard(w http.ResponseWriter, r *http.Request) {
	// Get the userID
	user := r.Header.Get("userid")
	userID, err := strconv.Atoi(user)
	if err != nil {
		http.Error(w, "The userID was unreadable", http.StatusInternalServerError)
		logging.Log("joinBoard", fmt.Sprintf("Userid[%v] could not be converted", user), err)
		return
	}

	params := mux.Vars(r)
	codeStr := params["code"]
	if codeStr == "" {
		http.Error(w, "Code empty", http.StatusBadRequest)
		return
	}

	db.JoinBoard(userID, codeStr)
}

func leaveBoard(w http.ResponseWriter, r *http.Request) {
	// todo
}

func createBoard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	codeStr := params["code"]
	if codeStr == "" {
		http.Error(w, "Code empty", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("could not read body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logging.Log("createBoard", err.Error(), err)
		return
	}

	db.InsertBoard(codeStr, string(body))
}

func getBoard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	codeStr := params["code"]
	if codeStr == "" {
		http.Error(w, "Code empty", http.StatusBadRequest)
		return
	}

	now := time.Now()
	day := now.Day()
	month := int(now.Month())
	year := now.Year()

	err := json.NewEncoder(w).Encode(db.GetBoard(codeStr, fmt.Sprintf("%d-%d-%d", day, month, year)))
	if err != nil {
		err = fmt.Errorf("could not convert to json")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logging.Log("boardsDisplayHandler", err.Error(), err)
		return
	}
}
