package controller

import (
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
	var query string
	var res model.ResponseResult

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Panic(err)
	}

	stmtOut, err := db.Prepare("SELECT username FROM users WHERE username  = ? ")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
	err = stmtOut.QueryRow(user.Username).Scan(&query)
	if err != nil {
		stmt, err := db.Prepare("INSERT INTO users(Username,FirstName,LastName,Password,Gender,Country,Age,Email) VALUES(?)")
		if err != nil {
			log.Panic(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(user.Username, user.FirstName, user.LastName, user.Password, user.Gender, user.Country, user.Age, user.Email)
		if err != nil {
			log.Panic(err)
		}
		res.Result = "Registration Successful"
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Error = "Already Registered Username"
	json.NewEncoder(w).Encode(res)

}

//HandleSignIn func
func HandleSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	var query model.User
	var res model.ResponseResult

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Panic(err)
	}

	stmtOut, err := db.Prepare("SELECT username FROM users WHERE (username  = ? OR email = ?) AND password =?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
	err = stmtOut.QueryRow(user.Email, user.Email, user.Password).Scan(&query)
	if err != nil {
		res.Error = "Invalid Email/Password"
		json.NewEncoder(w).Encode(res)
		return
	}
	res.Result = "Login Successful"
	json.NewEncoder(w).Encode(res)

}
