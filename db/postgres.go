package db

import (
	"fmt"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

// Connect postgres database
func Connect(dbName, dbHost, dbUser, dbPassword string) {
	var settings = postgresql.ConnectionURL{
		Database: dbName,
		Host:     dbHost,
		User:     dbUser,
		Password: dbPassword,
	}
	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}

	defer sess.Close()
	fmt.Printf("🚀 Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)

}
