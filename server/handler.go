package server

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *Server) handleRedirect(url string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, url, http.StatusTemporaryRedirect)
	}
}

func (s *Server) handleGetHealth() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.WithContext(request.Context()).WithField("URL", request.URL.String()).Debug("handling request")
		writer.WriteHeader(http.StatusOK)
	}
}
