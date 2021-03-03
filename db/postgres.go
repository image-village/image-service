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

	fmt.Printf("Collections in database %q:\n", sess.Name())

	// The Collections method returns references to all the collections in the
	// database.
	collections, err := sess.Collections()
	// get specific colelction by name
	// collection = sess.Collection("images") // TODO: get images collection
	if err != nil {
		log.Fatal("Collections: ", err)
	}

	for i := range collections {
		// Name returns the name of the collection.
		fmt.Printf("-> %q\n", collections[i].Name())
	}

	fmt.Printf("🚀 Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)

}
