package main

import (
	"dashboard/database"
	"dashboard/handlers"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	database.InitialiseDb()

	var dir string
	flag.StringVar(&dir, "dir", "./ui/build", "server for UI")
	flag.Parse()
	r.PathPrefix("/home").Handler(http.StripPrefix("/home", http.FileServer(http.Dir(dir))))

	r.HandleFunc("/signup", handlers.HandleSignUp).Methods("POST")
	r.HandleFunc("/signin", handlers.HandleSignIn).Methods("POST")
	r.HandleFunc("/editprofile", handlers.HandleEditProfile).Methods("PUT")

	log.Println("Started server on - 127.0.0.1:8000")
	go func() {
		log.Fatal(http.ListenAndServe(":8000", r))
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGINT)
	<-stop
	log.Println("Interrupt Encountered,..Closing Connections...")
	database.CloseDb()
}
