package config

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
)

/*
 * get mongodb connection
 */
func GetMongoDb() *mgo.Database {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")},
		Database: "users",
		Username: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("Unable to connect to DB %s", err)

		return nil
	}

	return session.DB("users")
}
