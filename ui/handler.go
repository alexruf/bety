package ui

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (u *WebUi) handleGetRoot() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if _, err := writer.Write([]byte("Hello World!")); err != nil {
			log.WithError(err).Error("Something went wrong")
		}
	}
}
