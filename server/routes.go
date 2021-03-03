package server

import (
	"github.com/lagbana/images/server/handlers"
	"github.com/lagbana/images/server/middlewares"
)

func (s *Server) initializeRoutes() {
	s.Router.GET("/", handlers.Home)
	s.Router.POST("/images", middlewares.AuthGuard(handlers.UploadImage))
} 
