package server

import (
	"fmt"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/lagbana/images/data"
	"github.com/lagbana/images/server/middlewares"
)

// Server -
type Server struct {
	Router *httprouter.Router
}

var s = Server{}

// Run -
func Run() {
	s.Router = httprouter.New()
	// Connect to database
	db, err := data.ConnectDB()
	if  err != nil {
		log.Fatal(err)
	}
	s.initializeRoutes()
	// Check for authentication on all incoming requests
	m := middlewares.GlobalAuthCheck(s.Router)
	fmt.Printf("🚀Connected to %s Database and Listening on port 8080\n", db.Session.Name())
	log.Fatal(http.ListenAndServe(":8080", m))
}
