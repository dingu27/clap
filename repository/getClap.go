package repository

import (
	"encoding/json"
	"net/http"

	"github.com/dingu27/clap/crud"
	"github.com/dingu27/clap/db"
	"github.com/gorilla/mux"
)

//GetClaps ...
func GetClaps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := db.GetMongoCon()
	if err != nil {
		panic(err)
	}

	obj := crud.NewMongo(db, "claps")
	params := mux.Vars(r)

	clap, err1 := obj.FindByImageName(params["imgName"])
	if err1 != nil {
		json.NewEncoder(w).Encode(clap)
		return
	}

	json.NewEncoder(w).Encode(clap)
}
