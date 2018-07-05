package dbconnection

import (
	"os"

	"gopkg.in/mgo.v2"
)

func GetMongoCon() (*mgo.Database, error) {
	host := os.Getenv("MONGO_HOST")
	dbName := "rr_clap"

	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}
	db := session.DB(dbName)

	return db, nil
}
