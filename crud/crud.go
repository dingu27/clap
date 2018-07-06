package crud

import (
	"github.com/dingu27/clap/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Mongo ...
type Mongo struct {
	db         *mgo.Database
	collection string
}

//NewMongo ...
func NewMongo(db *mgo.Database, collection string) *Mongo {
	return &Mongo{
		db:         db,
		collection: collection,
	}

}

//Save ...
func (r *Mongo) Save(clap *model.Clap) error {
	clap.Claps = 1
	err := r.db.C(r.collection).Insert(clap)
	return err
}

//Update ...
func (r *Mongo) Update(clap *model.Clap) error {
	clap.Claps = clap.Claps + 1
	err := r.db.C(r.collection).Update(bson.M{"imagename": clap.ImageName}, clap)
	return err
}

//FindByImageName ...
func (r *Mongo) FindByImageName(imgName string) (*model.Clap, error) {
	var clap model.Clap

	err := r.db.C(r.collection).Find(bson.M{"imagename": imgName}).One(&clap)

	if err != nil {
		return nil, err
	}
	return &clap, nil
}

//FindAll ...
func (r *Mongo) FindAll() (*model.Claps, error) {
	var claps model.Claps

	err := r.db.C(r.collection).Find(bson.M{}).All(&claps)
	if err != nil {
		return nil, err
	}

	return &claps, nil
}
