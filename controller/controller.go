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
	stmt, err := db.Prepare("INSERT INTO users(Username,FirstName,LastName,Password,Gender,Country,Age,Email) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(user.Username, user.FirstName, user.LastName, user.Password, user.Gender, user.Country, user.Age, user.Email)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("new user is created.")
}

//HandleSignIn func
func HandleSignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
