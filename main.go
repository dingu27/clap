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
	router.HandleFunc("/api/claps/get", repository.GetClaps).Methods("POST")
	router.HandleFunc("/api/claps/post", repository.CreateClap).Methods("POST")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":80", handler))
}
