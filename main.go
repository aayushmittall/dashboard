package main

import (
	"dashboard/handlers"
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	var dir string
	flag.StringVar(&dir, "dir", "./ui/build", "server for UI")
	flag.Parse()
	r.PathPrefix("/home").Handler(http.StripPrefix("/home", http.FileServer(http.Dir(dir))))

	r.HandleFunc("/signup", handlers.HandleSignUp).Methods("POST")
	r.HandleFunc("/signin", handlers.HandleSignIn).Methods("POST")

	log.Printf("Server Started")
	log.Fatal(http.ListenAndServe(":8000", r))

}
