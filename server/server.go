package server

import (
	"context"
	"database/sql"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Database *sql.DB
	Router   *mux.Router
}

func New(database *sql.DB) Server {
	router := mux.NewRouter()
	return Server{
		Database: database,
		Router:   router,
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.WithContext(ctx).Debug("Shutting down server...")
	return nil
}
