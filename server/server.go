package server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/lagbana/images/config"
	"github.com/lagbana/images/db"
	"github.com/lagbana/images/server/middlewares"
	"log"
	"net/http"
)

// Server -
type Server struct {
	Router *httprouter.Router
}

var s = Server{}

// Run -
func Run() {
	s.Router = httprouter.New()
	env := config.EnvSetup()
	s.initializeRoutes()
	// Connect to database
	db.Connect(env.DbName, env.DbHost, env.DbUser, env.DbPassword)

	// Check for authentication on all incoming requests
	m := middlewares.GlobalAuthCheck(s.Router)
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(":8080", m))
}
