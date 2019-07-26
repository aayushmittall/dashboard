package database

import (
	"dashboard/model"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

//InitialiseDb to start db server
func InitialiseDb() {
	var err error
	db, err = sql.Open("mysql", "root:mysql1234@tcp(localhost:3306)/dashboard")
	if err != nil {
		log.Panic(err.Error())
	}
}

//CloseDb to close DB connection
func CloseDb() {
	db.Close()
}

//GetUserByUsername func for sign_up
func GetUserByUsername(user *model.UserProfile) (*model.UserProfile, error) {
	var err error
	stmtOut, err := db.Prepare("SELECT username FROM user_profile WHERE username  = ? ")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
	var username string
	err = stmtOut.QueryRow(user.Username).Scan(&username)
	return user, err
}

//InsertUserProfile func for sign_up
func InsertUserProfile(user *model.UserProfile) error {
	var err error
	_, err = GetUserByUsername(user)
	if err != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
		if err != nil {
			panic(err.Error())
		}
		user.Password = string(hash)

		stmt, err := db.Prepare("INSERT INTO user_profile(Username,FirstName,LastName,Password,Gender,Country,Age,Email) VALUES(?,?,?,?,?,?,?,?)")
		if err != nil {
			log.Panic(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(user.Username, user.FirstName, user.LastName, user.Password, user.Gender, user.Country, user.Age, user.Email)
		if err != nil {
			log.Panic(err)
		}
	}
	return err
}

//LoginUser func for sign_in
func LoginUser(user *model.UserProfile) string {
	var err error
	var res string

	stmtOut, err := db.Prepare("SELECT password FROM user_profile WHERE (username  = ? OR email = ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()
	var password string
	err = stmtOut.QueryRow(user.Email, user.Email).Scan(&password)
	if err != nil {
		res = "Invalid Email/Username"
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))

	if err != nil {
		res = "Invalid password"
	} else {
		res = "Login Successful"
	}
	return res
}
