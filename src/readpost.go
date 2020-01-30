package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v6"
)

func readPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postid := params["postid"]
	userid := params["userid"]

	minioClient, err := minio.New(config.DORegion, config.DOAccessKey, config.DOSecret, true)
	if err != nil {
		log("Problem accessing DO Spaces, " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{Status: 500, Message: "Internal Server Error"})
		return
	}

	err = minioClient.FGetObject(config.DOBucket, userid+"-"+postid+".json", "/tmp/"+userid+"-"+postid+".json", minio.GetObjectOptions{})
	if err != nil {
		log("Error retrieving object " + userid + "-" + postid + ".json" + " from spaces, " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{Status: 500, Message: "Internal Server Error"})
		return
	}

	content, err := ioutil.ReadFile("/tmp/" + userid + "-" + postid + ".json")
	if err != nil {
		log("Error reading tmp file /tmp/" + userid + "-" + postid + ".json" + ", " + err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{Status: 500, Message: "Internal Server Error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(struct {
		Status  int    `json:"status"`
		Message []byte `json:"message"`
	}{Status: 200, Message: content})
}
