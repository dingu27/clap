package main

import (
	"log"
	"net/http"

	"github.com/dingu27/clap/repository"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	// Init Router
	router := mux.NewRouter()

	// Create our route handlers
	router.HandleFunc("/api/claps/{imgName}", repository.GetClaps).Methods("GET")
	router.HandleFunc("/api/claps", repository.CreateClap).Methods("POST")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":80", handler))
}
