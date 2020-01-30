package main

import "os"

type cfg struct {
	Listen      string
	APIKey      string
	DORegion    string
	DOAccessKey string
	DOSecret    string
	DOBucket    string
}

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
