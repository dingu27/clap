package db

import (
	"os"

	"gopkg.in/mgo.v2"
)

const database = "rr_clap"

// GetMongoCon database connection method
func GetMongoCon() (*mgo.Database, error) {
	host := os.Getenv("MONGO_HOST")
	dbName := database

	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}
	db := session.DB(dbName)

	return db, nil
}
