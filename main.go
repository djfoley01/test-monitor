package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/djfoley01/test-monitor/config"
	. "github.com/djfoley01/test-monitor/database"
	. "github.com/djfoley01/test-monitor/models"
)

var config = Config{}
var database = CDatabase{}

// GET list of clusters
func AllClustersEndPoint(w http.ResponseWriter, r *http.Request) {
	clusters, err := database.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, clusters)
}

// GET a cluster by its ID
func FindClusterEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cluster, err := database.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Cluster ID")
		return
	}
	respondWithJson(w, http.StatusOK, cluster)
}

// POST a new cluster
func CreateClusterEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var cluster Cluster
	if err := json.NewDecoder(r.Body).Decode(&cluster); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	cluster.ID = bson.NewObjectId()
	if err := database.Insert(cluster); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, cluster)
}

// PUT update an existing cluster
func UpdateClusterEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var cluster Cluster
	if err := json.NewDecoder(r.Body).Decode(&cluster); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := database.Update(cluster); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing cluster
func DeleteClusterEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var cluster Cluster
	if err := json.NewDecoder(r.Body).Decode(&cluster); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := database.Delete(cluster); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	database.Server = config.Server
	database.Database = config.Database
	database.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cluster", AllClustersEndPoint).Methods("GET")
	r.HandleFunc("/cluster", CreateClusterEndPoint).Methods("POST")
	r.HandleFunc("/cluster", UpdateClusterEndPoint).Methods("PUT")
	r.HandleFunc("/cluster", DeleteClusterEndPoint).Methods("DELETE")
	r.HandleFunc("/cluster/{id}", FindClusterEndpoint).Methods("GET")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
