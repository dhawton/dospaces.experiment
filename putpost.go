package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v6"
)

type postBody struct {
	data string
}

func putPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userid := params["userid"]
	postid := strconv.FormatInt(time.Now().Unix(), 16)

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

	_, err = minioClient.PutObject(config.DOBucket, userid+"-"+postid+".json", r.Body, r.ContentLength, minio.PutObjectOptions{})
	if err != nil {
		log("Error uploading object " + userid + "-" + postid + ".json" + " to spaces, " + err.Error())
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
		Status int    `json:"status"`
		PostID string `json:"postid"`
	}{Status: 200, PostID: postid})
}
