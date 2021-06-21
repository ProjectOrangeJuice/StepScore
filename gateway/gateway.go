package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"steps-gateway/db"

	"github.com/ProjectOrangeJuice/golib-logging/logging"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/auth/register", registerHandler)
	router.HandleFunc("/api/auth/login", accessHandler)
	log.Fatal(http.ListenAndServe(":9091", router))
}

func accessHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("username")
	id := db.GetUserID(user)
	if id == -1 {
		http.Error(w, "user not found", http.StatusUnauthorized)
		return
	}
	w.Header().Set("userid", strconv.Itoa(id))
	log.Printf("user auth -> %s", user)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	userB, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = fmt.Errorf("could not read body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logging.Log("registerHandler", err.Error(), err)
		return
	}
	user := string(userB)
	if user == "" {
		http.Error(w, "empty username", http.StatusBadRequest)
	}
	id := db.GetUserID(user)
	if id != -1 {
		http.Error(w, "user already exists", http.StatusConflict)
		return
	}

	db.AddUser(user)
	id = db.GetUserID(user)
	w.Header().Set("userid", strconv.Itoa(id))
	log.Printf("New user created -> %s", user)
}
