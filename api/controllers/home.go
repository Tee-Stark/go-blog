package controllers

import (
	"net/http"

	"github.com/tee-stark/go-blog/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to my blog app API")
}
