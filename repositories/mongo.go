package repositories

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

const (
	host = "localhost:27017"
)

// GetMongoSession retuns a session to the Mongo database
func GetMongoSession() (*mgo.Session, error) {
	log.Println("Attempting to connect to database...")

	session, err := mgo.DialWithInfo(
		&mgo.DialInfo{
			Addrs:   []string{host},
			Timeout: 10 * time.Second,
		},
	)
	if err != nil {
		log.Printf("Cannot connect to database: %s\n", err.Error())
	} else {
		log.Println("...connected to database.")
	}

	return session, err
}
