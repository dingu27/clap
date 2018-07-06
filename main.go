package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dingu27/clap/db"
	"github.com/dingu27/clap/model"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var claps model.Clap

func getClaps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := db.GetMongoCon()
	if err != nil {
		panic(err)
	}
	claps := model.Claps
	cl := db.C("image_claps_test").Find(bson.M{})
	json.NewEncoder(w).Encode(cl)
}

// func getBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	for _, item := range books {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Book{})
// }

func createClap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var clap model.Clap
	_ = json.NewDecoder(r.Body).Decode(&clap)
	clap.Claps = 1
	db, err := db.GetMongoCon()
	if err != nil {
		panic(err)
	}
	db.C("image_claps_test").Insert(clap)
	json.NewEncoder(w).Encode(clap)
}

// func updateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			var book Book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = params["id"]
// 			books = append(books, book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

// func deleteBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

func main() {

	// Init Router
	router := mux.NewRouter()

	// books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	// books = append(books, Book{ID: "2", Isbn: "454555", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Create our route handlers
	router.HandleFunc("/api/claps", getClaps).Methods("GET")
	// router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/claps", createClap).Methods("POST")
	// router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	// router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
