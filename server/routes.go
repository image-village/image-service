package server

import (
	"github.com/lagbana/images/server/handlers"
	"github.com/lagbana/images/server/middlewares"
)

func (s *Server) initializeRoutes() {
	s.Router.GET("/", middlewares.SetMiddlewareJSON(handlers.Home))
	s.Router.POST("/image", middlewares.SetMiddlewareJSON(handlers.UploadImage))
	
}