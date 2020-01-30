package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var config *cfg

func main() {
	gotenv.Load()

	log("Loading configuration variables")
	loadConfig()

	log("Setting up router")
	router := mux.NewRouter()

	log("Setting up routes and middleware")
	router.Use(authMiddleware)
	router.HandleFunc("/ping", ping).Methods("GET")
	router.HandleFunc("/post/{userid}", putPost).Methods("POST")
	router.HandleFunc("/post/{userid}/{postid}", readPost).Methods("GET")

	srv := http.Server{
		Handler:      router,
		Addr:         config.Listen,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log("Starting up server, listening on: " + config.Listen)
	err := srv.ListenAndServe()
	if err != nil {
		log("Error starting up server, " + err.Error())
		os.Exit(1)
	}
}
