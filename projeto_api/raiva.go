package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Client struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Clients []Client

var clients Clients

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/clients", GetClients).Methods("GET")
	router.HandleFunc("/clients", CreateClient).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetClients(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(clients)
}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	var newClient Client
	_ = json.NewDecoder(r.Body).Decode(&newClient)

	clients = append(clients, newClient)
	json.NewEncoder(w).Encode(newClient)
}
