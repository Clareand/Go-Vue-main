package controllers

import (
	"github.com/Clareand/rest-api/api/middlewares"
)

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	s.Router.HandleFunc("/api/v1/auth/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/api/v1/auth/register", middlewares.SetMiddlewareJSON(s.Register)).Methods("POST")
}
