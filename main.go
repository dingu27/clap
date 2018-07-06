package main

import (
	"log"
	"net/http"

	"github.com/dingu27/clap/repository"
	"github.com/gorilla/mux"
)

func main() {

	// Init Router
	router := mux.NewRouter()

	// Create our route handlers
	router.HandleFunc("/api/claps/{imgName}", repository.GetClaps).Methods("GET")
	router.HandleFunc("/api/claps", repository.CreateClap).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
