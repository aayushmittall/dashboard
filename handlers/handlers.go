package handlers

import (
	"dashboard/database"
	"dashboard/model"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var db *sql.DB
var err error

//HandleSignUp func
func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	var res string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Panic(err)
	}

	res = database.InsertUser(&user)
	json.NewEncoder(w).Encode(res)

}

//HandleSignIn func
func HandleSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	var res string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Panic(err)
	}

	res = database.LoginUser(&user)
	json.NewEncoder(w).Encode(res)

}
