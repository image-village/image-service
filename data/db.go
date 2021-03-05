package data

import (
	"github.com/lagbana/images/config"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

// Database holds database prosgres table structure
type Database struct {
	db.Session
	Images *ImagesStore
}

// DB is a pointer to the database interface
var DB *Database

// ConnectDB initializes / opens a new database connection
func ConnectDB() (*Database, error) {
	env := config.EnvSetup()
	var settings = postgresql.ConnectionURL{
		Database: env.DbName,
		Host:     env.DbHost,
		User:     env.DbUser,
		Password: env.DbPassword,
	}
	db := &Database{}
	sess, err := postgresql.Open(settings)
	if err != nil {
		return nil, err
	}
	db.Session = sess
	db.Images = Images(db.Session)
	// global instance for access across modules
	DB = db

	return db, nil
}
