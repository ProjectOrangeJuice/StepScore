package main

import (
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

	router.HandleFunc("/steps/{day}/{month}/{year}", stepsHandler).Methods("POST")
	router.HandleFunc("/board/{code}", boardHandler).Methods("POST", "GET", "DELETE", "PUT")
	router.HandleFunc("/boards", boardsDisplayHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":9090", router))
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

}

func boardsDisplayHandler(w http.ResponseWriter, r *http.Request) {

}
