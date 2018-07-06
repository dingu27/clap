package repository

import (
	"encoding/json"
	"net/http"

	"github.com/dingu27/clap/db"
	"github.com/dingu27/clap/model"
	"gopkg.in/mgo.v2/bson"
)

//CreateClap ...
func CreateClap(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var clap model.Clap
	_ = json.NewDecoder(r.Body).Decode(&clap)
	db, err := db.GetMongoCon()
	if err != nil {
		panic(err)
	}

	var claps model.Claps
	err1 := db.C("image_claps_test").Find(bson.M{}).All(&claps)
	if err1 != nil {
		panic(err1)
	}

	for _, item := range claps {
		if item.ImageName == clap.ImageName {
			item.Claps = item.Claps + 1
			err2 := db.C("image_claps_test").Update(bson.M{"imagename": item.ImageName}, item)
			if err2 != nil {
				panic(err2)
			}
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	clap.Claps = 1
	db.C("image_claps_test").Insert(clap)
	json.NewEncoder(w).Encode(clap)
}
