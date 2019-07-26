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

//HandleSignUp func
func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var err error

	var user model.UserProfile
	var res string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Panic(err)
	}

	err = database.InsertUserProfile(&user)
	if err == nil {
		res = "Sign Up Successful"
		json.NewEncoder(w).Encode(res)

	}

}

//HandleSignIn func
func HandleSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.UserProfile
	var res string
	var err error

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
