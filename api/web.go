package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ProajgeOrangeJuice/StepCounter/db"
	"github.com/ProjectOrangeJuice/golib-logging/logging"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/steps/{day}/{month}/{year}", stepsHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/board/{code}", boardHandler).Methods("POST", "GET", "DELETE", "PUT", "OPTIONS")
	router.HandleFunc("/api/boards", boardsDisplayHandler).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":9090", router))
	//log.Fatal(http.ListenAndServe(":9090", handlers.CompressHandler(cors(router))))
}

func stepsHandler(w http.ResponseWriter, r *http.Request) {
	// Get the userID
	user := r.Header.Get("userid")
	userID, err := strconv.Atoi(user)
	if err != nil {
		http.Error(w, "The userID was unreadable", http.StatusInternalServerError)
		logging.Log("stepsHandler", fmt.Sprintf("Userid[%v] could not be converted", user), err)
		return
	}
	params := mux.Vars(r)

	day := params["day"]
	month := params["month"]
	year := params["year"]

	now := time.Now()

	// Validate that the date is today
	if day != strconv.Itoa(now.Day()) || month != strconv.Itoa(int(now.Month())) || year != strconv.Itoa(now.Year()) {
		http.Error(w, "Date should be today", http.StatusBadRequest)
		return
	}

	stepsStr := r.FormValue("steps")
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		http.Error(w, "Could not read steps", http.StatusBadRequest)
		return
	}

	// Set the steps for the date
	db.InsertSteps(userID, steps, fmt.Sprintf("%s-%s-%s", day, month, year))

}

func boardHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createBoard(w, r)
	case http.MethodGet:
		getBoard(w, r)
	case http.MethodDelete:
		leaveBoard(w, r)
	case http.MethodPut:
		joinBoard(w, r)
	}
}

func boardsDisplayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	// Get the userID
	user := r.Header.Get("userid")
	userID, err := strconv.Atoi(user)
	if err != nil {
		http.Error(w, "The userID was unreadable", http.StatusInternalServerError)
		logging.Log("boardsDisplayHandler", fmt.Sprintf("Userid[%v] could not be converted", user), err)
		return
	}

	codes := db.GetBoards(userID)

	err = json.NewEncoder(w).Encode(codes)
	if err != nil {
		err = fmt.Errorf("could not convert to json")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logging.Log("boardsDisplayHandler", err.Error(), err)
		return
	}

}
