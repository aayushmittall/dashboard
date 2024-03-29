package database

import (
	"dashboard/model"
	"database/sql"
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

//InitialiseDb to start db server
func InitialiseDb() error {
	var err error
	db, err = sql.Open("mysql", "root:mysql1234@tcp(localhost:3306)/dashboard")
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

//CloseDb to close DB connection
func CloseDb() {
	db.Close()
}

//GetUserByUsername func to get user by username
func GetUserByUsername(username string) (*model.UserProfile, error) {
	var err error
	var profile *model.UserProfile
	stmtOut, err := db.Prepare("SELECT * FROM user_profile WHERE Username  = ? or Email=? ")
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer stmtOut.Close()
	err = stmtOut.QueryRow(username).Scan(&profile.ID, &profile.Username, &profile.FirstName, &profile.LastName, &profile.Password, &profile.Gender, &profile.Country, &profile.Age, &profile.Email)
	return profile, err
}

//EncryptPassword to secure passwords
func EncryptPassword(userpassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(userpassword), 4)
	if err != nil {
		log.Print(err)
		return userpassword, err
	}
	userpassword = string(hash)
	return userpassword, nil
}

//InsertUserProfile func to insert user profile
func InsertUserProfile(user *model.UserProfile) error {
	var profile *model.UserProfile
	var hashedpassword string
	var err error
	profile, err = GetUserByUsername(user.Username)
	if profile == nil && err != nil {
		hashedpassword, err = EncryptPassword(user.Password)
		if err != nil {
			log.Print(err)
			return err
		}
		user.Password = hashedpassword
		stmt, err := db.Prepare("INSERT INTO user_profile(Username,FirstName,LastName,Password,Gender,Country,Age,Email) VALUES(?,?,?,?,?,?,?,?)")
		if err != nil {
			log.Print(err)
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(user.Username, user.FirstName, user.LastName, user.Password, user.Gender, user.Country, user.Age, user.Email)
		if err != nil {
			log.Print(err)
			return err
		}
	}
	return nil
}

//InsertUserAuth to insert token in db.
func InsertUserAuth(UserID int, Token string) error {
	var userAuth *model.UserAuth
	userAuth.UserID = UserID
	userAuth.Token = Token
	userAuth.TimeGenerated = time.Now()
	stmt, err := db.Prepare("INSERT INTO user_auth(UserID,Token,TimeGenerated) VALUES(?,?,?)")
	if err != nil {
		log.Print(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userAuth.UserID, userAuth.Token, userAuth.TimeGenerated)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil

}

//GenerateToken func to generate random token string
func GenerateToken() string {
	bytes := make([]byte, 25)
	for i := 0; i < 25; i++ {
		bytes[i] = byte(65 + rand.Intn(25)) //A=65 and Z = 65+25
	}
	return string(bytes)
}

//LoginUser func for sign_in
func LoginUser(user *model.UserProfile) (string, error) {
	var profile *model.UserProfile
	var token string
	var err error
	profile, err = GetUserByUsername(user.Username)
	if err != nil {
		log.Print(err)
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(profile.Password), []byte(user.Password))
	if err != nil {
		log.Print(err)
		return "", err
	}
	token = GenerateToken()
	err = InsertUserAuth(profile.ID, token)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return token, nil
}
