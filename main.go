package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

type cfg struct {
	Listen      string
	APIKey      string
	DORegion    string
	DOAccessKey string
	DOSecret    string
	DOBucket    string
}

var config *cfg

func loadConfig() {
	var val string
	var ok bool

	if config == nil {
		config = &cfg{}
	}

	if val, ok = os.LookupEnv("PORT"); ok {
		config.Listen = ":" + val
	} else {
		config.Listen = ":1776"
	}

	if val, ok = os.LookupEnv("API_KEY"); ok {
		config.APIKey = val
	} else {
		log("API_KEY not set. Cannot continue.")
		os.Exit(1)
	}

	if val, ok = os.LookupEnv("DO_REGION"); ok {
		config.DORegion = val
	} else {
		log("DO Region not set. Cannot continue.")
		os.Exit(1)
	}

	if val, ok = os.LookupEnv("DO_KEY"); ok {
		config.DOAccessKey = val
	} else {
		log("DO Access Key not set. Cannot continue.")
		os.Exit(1)
	}

	if val, ok = os.LookupEnv("DO_SECRET"); ok {
		config.DOSecret = val
	} else {
		log("DO Secret not set. Cannot continue.")
		os.Exit(1)
	}

	if val, ok = os.LookupEnv("DO_BUCKET"); ok {
		config.DOBucket = val
	} else {
		log("DO Bucket not set. Cannot continue.")
		os.Exit(1)
	}
}

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
