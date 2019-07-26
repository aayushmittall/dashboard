package database

import (
	"dashboard/model"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

//InitialiseDb to start db server
func InitialiseDb() error {
	var err error
	db, err = sql.Open("mysql", "root:mysql1234@tcp(localhost:3306)/dashboard")
	return err
}

//CloseDb to close DB connection
func CloseDb() {
	db.Close()
}

//GetUserByUsername func to get user by username
func GetUserByUsername(username string) (*model.UserProfile, error) {
	var err error
	var profile *model.UserProfile
	stmtOut, err := db.Prepare("SELECT * FROM user_profile WHERE username  = ? ")
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()
	err = stmtOut.QueryRow(username).Scan(&profile.Username, &profile.FirstName, &profile.LastName, &profile.Password, &profile.Gender, &profile.Country, &profile.Age, &profile.Email)
	return profile, err
}

//EncryptPassword to secure passwords
func EncryptPassword(userpassword string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(userpassword), 4)
	if err != nil {
		return err.Error()
	}
	userpassword = string(hash)
	return userpassword
}

//InsertUserProfile func to insert user profile
func InsertUserProfile(user *model.UserProfile) error {
	var profile *model.UserProfile
	var err error
	profile, err = GetUserByUsername(user.Username)
	if profile == nil && err != nil {
		user.Password = EncryptPassword(user.Password)
		stmt, err := db.Prepare("INSERT INTO user_profile(Username,FirstName,LastName,Password,Gender,Country,Age,Email) VALUES(?,?,?,?,?,?,?,?)")
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(user.Username, user.FirstName, user.LastName, user.Password, user.Gender, user.Country, user.Age, user.Email)
		if err != nil {
			return err
		}
	}
	return err
}

//GenerateToken func to add token in db
func GenerateToken(username string) error {
	var userAuth *model.UserAuth
	var err error
	userAuth.Token = EncryptPassword(username)
	userAuth.Username = username
	stmt, err := db.Prepare("INSERT INTO user_auth(Username,Token) VALUES(?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userAuth.Username, userAuth.Token)
	if err != nil {
		return err
	}
	return nil
}

//LoginUser func for sign_in
func LoginUser(user *model.UserProfile) error {
	var profile *model.UserProfile
	var err error
	profile, err = GetUserByUsername(user.Username)
	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(user.Password))
		if err != nil {
			return err
		}
		err = GenerateToken(user.Username)
	}
	return nil
}
