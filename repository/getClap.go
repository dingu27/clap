package repository

import (
	"encoding/json"
	"net/http"

	"github.com/dingu27/clap/crud"
	"github.com/dingu27/clap/db"
	"github.com/dingu27/clap/model"
)

//GetClaps ...
func GetClaps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := db.GetMongoCon()
	if err != nil {
		panic(err)
	}

	obj := crud.NewMongo(db, "claps")

	var clap model.Clap
	_ = json.NewDecoder(r.Body).Decode(&clap)

	filteredClap, err1 := obj.FindByImageName(clap.ImageName)
	if err1 != nil {
		json.NewEncoder(w).Encode(filteredClap)
		return
	}

	json.NewEncoder(w).Encode(filteredClap)
}
