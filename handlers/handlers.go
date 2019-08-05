package handlers

import (
	"dashboard/database"
	"dashboard/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Response struct
type Response struct {
	Token  string `json:"token"`
	Result string `json:"result"`
}

//HandleSignUp func
func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.UserProfile
	var response Response
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

	err = database.InsertUserProfile(&user)
	if err != nil {
		log.Print(err)
		json.NewEncoder(w).Encode(err.Error())
	}
	response.Result = "Sign Up Successful"
	json.NewEncoder(w).Encode(response)

}

//HandleSignIn func
func HandleSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.UserProfile
	var response Response
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

	response.Token, err = database.LoginUser(&user)
	if err != nil {
		log.Print(err)
		json.NewEncoder(w).Encode(err.Error())
	}
	response.Result = "Login Successful"
	json.NewEncoder(w).Encode(response)

}
