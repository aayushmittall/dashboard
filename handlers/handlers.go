package handlers

import (
	"dashboard/database"
	"dashboard/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//HandleSignUp func
func HandleSignUp(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	var err error

	var user model.UserProfile
	var res string

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	err = database.InsertUserProfile(&user)
	if err == nil {
		res = "Sign Up Successful"
		json.NewEncoder(w).Encode(res)

	}
	return err

}

//HandleSignIn func
func HandleSignIn(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	var user model.UserProfile
	var res string
	var err error

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return err
	}

	res = database.LoginUser(&user)
	json.NewEncoder(w).Encode(res)
	return err
}
