package repository

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dingu27/clap/crud"
	"github.com/dingu27/clap/db"
	"github.com/dingu27/clap/model"
)

//CreateClap ...
func CreateClap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := db.GetMongoCon()
	if err != nil {
		panic(err)
	}

	obj := crud.NewMongo(db, "claps")

	var tClap model.Clap
	_ = json.NewDecoder(r.Body).Decode(&tClap)

	clap, err1 := obj.FindByImageName(tClap.ImageName)
	if err1 != nil {
		err2 := obj.Save(&tClap)
		if err2 != nil {
			log.Fatal(err2)
		}
		json.NewEncoder(w).Encode(tClap)
		return
	}

	err3 := obj.Update(clap)
	if err3 != nil {
		log.Fatal(err3)
	}
	json.NewEncoder(w).Encode(clap)
}
