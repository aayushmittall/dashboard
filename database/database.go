package database

import (
	"dashboard/model"
	"database/sql"
	"log"
)

var db *sql.DB
var err error

//Getdb func
func Getdb() {
	db, err = sql.Open("mysql", "root:mysql1234@tcp(localhost:3306)/dashboard")
	if err != nil {
		log.Panic(err.Error())
		return
	}
	defer db.Close()
}

//InsertUser func for sign_up
func InsertUser(user *model.User) string {
	stmtOut, err := db.Prepare("SELECT username FROM users WHERE username  = ? ")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
	var query string
	var result string
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
		result = "Registration Successful"
	}
	result = "Already Registered Username"
	return result
}

//LoginUser func for sign_in
func LoginUser(user *model.User) string {
	var res string

	stmtOut, err := db.Prepare("SELECT username FROM users WHERE (username  = ? OR email = ?) AND password =?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
	var query string
	err = stmtOut.QueryRow(user.Email, user.Email, user.Password).Scan(&query)
	if err != nil {
		res = "Invalid Email/Password"
	}
	res = "Login Successful"
	return res
}
