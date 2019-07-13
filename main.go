package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//===========================================================================================
//===========================================================================================
//MODEL
var db *sql.DB
var err error

/*create table users(Username varchar(10) primary,
FirstName varchar(10),LastName varchar(10),Password varchar(20),
Gender varchar(10),Country varchar(20),Age varchar(10)); */

//User struct
type User struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Country   string `json:"country"`
	Age       string `json:"age"`
	Email     string `json:"email"`
}

//===========================================================================================
//CONTROLLERSb
func handleSignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	stmt, err := db.Prepare("INSERT INTO users(Username,FirstName,LastName,Password,Gender,Country,Age,Email) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// keyVal := make(map[string]string)
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}
	// Username := keyVal["username"]
	// FirstName := keyVal["firstname"]
	// LastName := keyVal["lastname"]
	// Password := keyVal["password"]
	// Gender := keyVal["gender"]
	// Country := keyVal["country"]
	// Age := keyVal["age"]
	// Email := keyVal["email"]

	_, err = stmt.Exec(user.Username, user.FirstName, user.LastName, user.Password, user.Gender, user.Country, user.Age, user.Email)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("new user is created.")
}

// func handleSignIn(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	keyVal := make(map[string]string)
// 	json.Unmarshal(body, &keyVal)
// 	Username := keyVal["username"]
// 	Password := keyVal["password"]
// 	_, err := db.Prepare("select * from users where Username=? and Password=?", Username, Password)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// }

func main() {
	//===========================================================================================

	r := mux.NewRouter()
	r.HandleFunc("/signup", handleSignUp)
	// r.HandleFunc("/SignIn", handleSignUp)

	//===========================================================================================

	var dir string
	flag.StringVar(&dir, "dir", "./build", "build")
	flag.Parse()

	r.PathPrefix("/home").Handler(http.StripPrefix("/home", http.FileServer(http.Dir(dir))))
	//===========================================================================================
	db, err = sql.Open("mysql", "root:mysql1234@tcp(127.0.0.1:3306)/dashboard")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//===========================================================================================

	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:8000",
	}
	fmt.Printf("Server Started")
	log.Fatal(srv.ListenAndServe())
	//===========================================================================================

}
