package main

import (
	"dashboard/controller"
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

func main() {

	r := mux.NewRouter()

	db, err = sql.Open("mysql", "root:mysql1234@tcp(localhost:3306)/dashboard")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	var dir string
	flag.StringVar(&dir, "dir", "./ui/build", "server for UI")
	flag.Parse()
	r.PathPrefix("/home").Handler(http.StripPrefix("/home", http.FileServer(http.Dir(dir))))

	r.HandleFunc("/signup", controller.HandleSignUp)
	r.HandleFunc("/signin", controller.HandleSignIn)

	log.Printf("Server Started")
	http.ListenAndServe(":8000", r)

}
