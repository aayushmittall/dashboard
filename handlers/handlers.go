package handlers

import (
	"dashboard/database"
	"dashboard/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//HandleSignUp func
func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.UserProfile
	var err error
	var res string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		json.NewEncoder(w).Encode(err.Error())
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Print(err)
		json.NewEncoder(w).Encode(err.Error())
	}

	err = database.InsertUserProfile(&user)
	if err != nil {
		log.Print(err)
		json.NewEncoder(w).Encode(err.Error())
	}
	res = "Sign Up Successful"
	json.NewEncoder(w).Encode(res)

}

//HandleSignIn func
func HandleSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.UserProfile
	var token string
	var err error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		json.NewEncoder(w).Encode(err.Error())

	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Print(err)
		json.NewEncoder(w).Encode(err.Error())

	}

	token, err = database.LoginUser(&user)
	if err != nil {
		log.Print(err)
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(token)

}
