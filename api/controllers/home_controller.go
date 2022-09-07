package controllers

import (
	"net/http"

	"github.com/Clareand/rest-api/api/helpers"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	helpers.JSON(w, http.StatusOK, "API fine")
}
