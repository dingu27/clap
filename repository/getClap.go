package repository

import (
	"encoding/json"
	"net/http"

	"github.com/dingu27/clap/db"
	"github.com/dingu27/clap/model"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

//GetClaps ...
func GetClaps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := db.GetMongoCon()
	if err != nil {
		panic(err)
	}
	var claps model.Claps
	err1 := db.C("image_claps_test").Find(bson.M{}).All(&claps)
	if err1 != nil {
		panic(err1)
	}

	params := mux.Vars(r)

	var filterClap model.Claps

	for _, item := range claps {
		if item.ImageName == params["imgName"] {
			filterClap = append(filterClap, item)
		}
	}
	json.NewEncoder(w).Encode(filterClap)
}
